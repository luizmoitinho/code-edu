package domain

import (
	"fmt"
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

//User ...
type User struct {
	Base     `valid:"required"`
	Name     string `json:"name" gorm:"type:varchar(255)" valid:"notnull"`
	Email    string `json:"email" gorm:"type:varchar(255);unique_index" valid:"notnull, email"`
	Password string `json:"-" gorm:"type:varchar(255)" valid:"notnull"`
	Token    string `json:"token" gorm:"type:varchar(255);unique_index" valid:"notnull, uuid"`
}

func init() {

	govalidator.SetFieldsRequiredByDefault(true)

}

//NewUser ... Validate and create a new user instance
func NewUser(name string, email string, password string) (*User, error) {

	user := &User{
		Name:     name,
		Email:    email,
		Password: password,
	}

	err := user.Prepare()

	if err != nil {
		return nil, err
	}

	return user, nil

}

func (user *User) validate() error {
	_, err := govalidator.ValidateStruct(user)

	if err != nil {
		return err
	}

	return nil
}

//Prepare ... create a token and validate password
func (user *User) Prepare() error {

	if user.Password == "" {
		return fmt.Errorf("Password cannot be blank")
	}

	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Printf("error in user.Prepare(): %v", err)
		return err
	}

	user.ID = uuid.NewV4().String()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.Password = string(password)
	user.Token = uuid.NewV4().String()

	err = user.validate()

	if err != nil {
		fmt.Printf("error in user.Prepare(), user.validate(): %v", err)
		return err
	}

	return nil

}
