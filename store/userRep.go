package store 

import (
	"github.com/qwertyqq2/model"
)

type UserRepository struct{
	store *Store
}

func(r *UserRepository) Create(user *model.User) error{
	_, err:= r.store.db.Exec("insert into USERS values (?, ?)", 
	user.ID, user.Email); if err!=nil{
		return err
	}
	return nil
}

func(r *UserRepository) GetUId() uint {
	var id uint
	if err:=r.store.db.QueryRow("SELECT MAX(id) FROM USERS;").Scan(&id);err!=nil{
		return 0
	}
	return id
}

func(r *UserRepository) GetEmailById(id uint) string{
	var email string
	if err:=r.store.db.QueryRow("SELECT Email FROM USERS WHERE 	Id = ?;", id).Scan(&email);err!=nil{
		return ""	
	}
	return email
}
