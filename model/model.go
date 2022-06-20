package model

type User struct{
	ID uint `json: "id"`
	Email string `json: "email"`
}

type Transaction struct{
	ID uint `json: "id"`
	UID uint `json: "uid"`
	Email string	`json: "email"`
	Currency string	`json: "cur"`
	Sum uint	`json: "sum"`
	TimeOfCreation string	`json: "timeCreation"`
	TimeOfLastChange string	`json: "timeChange"`
	Status string	`json: "status"`
}
