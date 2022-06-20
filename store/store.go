package store

import(
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Store struct{
	config *Config
	db *sql.DB
	userRep *UserRepository
	txRep *txRepository

}

func New(config *Config) *Store{
	return &Store{
		config: config,		
	}
}

func (s *Store) Open() error{
	db, err := sql.Open("mysql", s.config.URL)
	if err!=nil{
		return err
	}
	if err:= db.Ping();err!=nil{
		return err
	}
	s.db = db
	return nil
}

func (s *Store) Close() {
	s.db.Close()
}

func (s *Store) User() *UserRepository {
	if s.userRep != nil {
		return s.userRep
	}
	s.userRep = &UserRepository{
		store: s,
	}
	return s.userRep
}

func (s *Store) Tx() *txRepository{
	if s.txRep != nil {
		return s.txRep
	}
	s.txRep = &txRepository{
		store: s,
	}
	return s.txRep
}
