package apiServer

import (
	"github.com/sirupsen/logrus"
	"github.com/gorilla/mux"
	"net/http"
	"github.com/qwertyqq2/store"
	"github.com/qwertyqq2/model"
	"encoding/json"
	"time"
	"fmt"
)

type APIServer struct{
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store *store.Store
}

func New(config *Config) *APIServer{
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (s *APIServer) Start() error{
	if err:=s.configuresLogger();err!=nil{
		return err
	}
	s.configuresRouter()
	if err:= s.configuresStore();err!=nil{
		return err
	}
	s.logger.Info("start!")
	return http.ListenAndServe(s.config.BindAddr, s.router)
}


func (s *APIServer) configuresStore() error{
	store_ := store.New(s.config.Store)
	if err:=store_.Open(); err!=nil{
		return err
	}
	s.store = store_
	return nil
}

func (s *APIServer) configuresLogger() error{
	level, err:= logrus.ParseLevel(s.config.LogLevel)
	if err!=nil{
		return err
	}
	s.logger.SetLevel(level)
	return err
}


func(s *APIServer) configuresRouter(){
	s.router.HandleFunc("/user", s.handleUsersCreate()).Methods("POST")	
	s.router.HandleFunc("/transaction", s.handleTXCreate()).Methods("POST")
	s.router.HandleFunc("/transactionsId", s.handleTxId()).Methods("POST")
	s.router.HandleFunc("/transactionsEmail", s.handleTxEmail()).Methods("POST")
} 





func (s *APIServer) handleUsersCreate() http.HandlerFunc {
	type request struct {
		Email   string `json:"email"`
	}	
	return func(w http.ResponseWriter, r *http.Request){
		req := &request{}
		if err:=json.NewDecoder(r.Body).Decode(req);  err!=nil{
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		id := s.store.User().GetUId() +1
		user := &model.User{
			ID: id,
			Email: req.Email,
		}
		if err:= s.store.User().Create(user); err!=nil{
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		s.respond(w, r, http.StatusCreated, user)

	}
}


func(s *APIServer) handleTXCreate() http.HandlerFunc{
	type request struct{
		UID uint `json: "uid"`
		Email string `json: "email"`
		Currency string `json: "cur"`
		Sum uint `json: "sum"`
	}
	return func(w http.ResponseWriter, r *http.Request){
		req := &request{}
		if err:=json.NewDecoder(r.Body).Decode(req);  err!=nil{
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		id := s.store.Tx().GetTId() +1 
		fmt.Println(id)

		timeCreation := time.Unix(time.Now().Unix(), 0).String()
		tx := &model.Transaction{
			ID: id,
			UID: req.UID,
			Email: req.Email,
			Currency: req.Currency,
			Sum: req.Sum,
			TimeOfCreation: timeCreation,
			TimeOfLastChange: timeCreation,
			Status: "New", 
		}
		if err:= s.store.Tx().Create(tx); err!=nil{
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		s.respond(w, r, http.StatusCreated, tx)
		
		////////Переход в статус Успех\Неудача//////////

		email:=s.store.User().GetEmailById(req.UID)
		status:=""
		if email == ""{
			s.respond(w, r, http.StatusBadRequest, "Not detected")
			status = "False"
		}else if email!=req.Email{
			s.respond(w, r, http.StatusBadRequest, "Not match")
			status = "False"
		}else{
			status = "True"
		}
		t:= time.Unix(time.Now().Unix(), 0).String()
		if err:= s.store.Tx().SetStatus(status, id, t); err!=nil{
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		s.respond(w, r, http.StatusCreated, tx)	
	}
}

///////Список платежей по Id/////////////
func(s *APIServer) handleTxId() http.HandlerFunc{ 
	type request struct{
		UID uint `json: "uid"`
	}
	return func(w http.ResponseWriter, r *http.Request){
		req := &request{}
		if err:=json.NewDecoder(r.Body).Decode(req);  err!=nil{
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		Tx:=s.store.Tx().GetAllTxById(req.UID)
		fmt.Println(Tx)
	}
}

///////Список платежей по Email/////////////
func(s *APIServer) handleTxEmail() http.HandlerFunc{ 
	type request struct{
		Email string `json: "email"`
	}
	return func(w http.ResponseWriter, r *http.Request){
		req := &request{}
		if err:=json.NewDecoder(r.Body).Decode(req);  err!=nil{
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		Tx:=s.store.Tx().GetAllTxByEmail(req.Email)
		fmt.Println(Tx)
	}
}


func (s *APIServer) error(w http.ResponseWriter, r *http.Request, code int, err error){
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *APIServer) respond(w http.ResponseWriter, r *http.Request, code int , data interface{}){
	w.WriteHeader(code)
	if data!=nil{
		json.NewEncoder(w).Encode(data)
	}
}
