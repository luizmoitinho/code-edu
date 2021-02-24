package repositories

import (
	"log"
	"projects/code-edu/domain"

	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	Insert(user *domain.User) (*domain.User, error)
	Update(user *domain.User) (*domain.User, error)
	Delete(user *domain.User) (bool, error)
}

type UserRepositoryDb struct {
	Db *gorm.DB
}

func (repo UserRepositoryDb) Insert(user *domain.User) (*domain.User, error) {

	err := user.Prepare()
	if err != nil {
		log.Fatalf("Error durigin the user valiation, user_repository.Insert(): %v", err)
		return nil, err
	}

	err = repo.Db.Create(user).Error
	if err != nil {
		log.Fatalf("Error to persist user, user_repository.Insert(): %v", err)
		return nil, err
	}
	return user, nil
}
