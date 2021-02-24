package domain

import (
	"log"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Base
	Name     string `json:"name" gorm:"type:varchar(255)"`
	Email    string `json:"email" gorm:"type:varchar(255);unique_index"`
	Password string `json:"-" gorm:"type:varchar(255)"`
	Token    string `json:"token" gorm:"type:varchar(255);unique_index"`
}

func NewUser() *User {
	return &User{}
}

//Prepare ... create a token and validate password
func (user *User) Prepare() error {

	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("error in user.Prepare(): %v", err)
		return err
	}

	user.Password = string(password)
	user.Token = uuid.NewV4().String()

	err = user.validate()
	if err != nil {
		log.Fatal("error in user.Prepare(), user.validate(): %v", err)
		return err
	}

	return nil

}

func (user *User) validate() error {
	return nil
}
