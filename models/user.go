package models

import (
	"github.com/nguyendangcuong201004/event-booking-api/db"
	"github.com/nguyendangcuong201004/event-booking-api/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (user User) Save() error {
	query := `
	INSERT INTO users (email, password)
	VALUES(?, ?)	
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	hashpassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(user.Email, hashpassword)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	user.ID = id
	return err
}
