package store 

import (
	"github.com/qwertyqq2/model"
	"fmt"
)

type txRepository struct{
	store *Store
}


func(r *txRepository) Create(tx *model.Transaction) error{
	_, err:= r.store.db.Exec("insert into TRANSACTIONS values (?, ?, ?, ?, ?, ? ,?)", 
	tx.ID, tx.UID, tx.Email, tx.Currency, tx.TimeOfCreation, tx.TimeOfLastChange, tx.Status); if err!=nil{
		return err
	}
	return nil
}

func(r *txRepository) GetTId() uint {
	var id uint
		if err:=r.store.db.QueryRow("SELECT MAX(Id) FROM TRANSACTIONS;").Scan(&id);err!=nil{
			return 0
	}
	return id
}


func(r *txRepository) SetStatus(s string, id uint, t string) error{
	_, err:= r.store.db.Exec("UPDATE TRANSACTIONS SET Status = ?, TimeOfLastChange = ? WHERE Id = ?",
		s, t, id); if err!=nil{
			return err
	} 
	return nil
}

func(r *txRepository) GetAllTxById(id uint) []*model.Transaction{
	txs:= make([]*model.Transaction, 0)
	rows, err:=r.store.db.Query("SELECT * FROM TRANSACTIONS WHERE IdUser = ? AND Status = ?;", id, "True")
	if err!=nil{
		fmt.Println(err)
	}
	for rows.Next() {
		tx:=new(model.Transaction)
        rows.Scan(&tx.ID, &tx.UID, &tx.Email, &tx.Currency, &tx.TimeOfCreation, &tx.TimeOfLastChange, &tx.Status)
        txs = append(txs, tx)
    }
	return txs
}

func(r *txRepository) GetAllTxByEmail(email string) []*model.Transaction{
	txs:= make([]*model.Transaction, 0)
	rows, err:=r.store.db.Query("SELECT * FROM TRANSACTIONS WHERE Email = ? AND Status = ?;", email, "True")
	if err!=nil{
		fmt.Println(err)
	}
	for rows.Next() {
		tx:=new(model.Transaction)
        rows.Scan(&tx.ID, &tx.UID, &tx.Email, &tx.Currency, &tx.TimeOfCreation, &tx.TimeOfLastChange, &tx.Status)
        txs = append(txs, tx)
    }
	return txs
}
