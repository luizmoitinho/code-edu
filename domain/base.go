package domain

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

//Base ...
type Base struct {
	ID        string    `gorm:"type:uuid;primary_key" valid:"uuid"`
	CreatedAt time.Time `valid:"-"`
	UpdatedAt time.Time `valid:"-"`
	//DeletedAt time.Time `json:"deleted_at"`
}

//BeforeCreate ...
func (base *Base) BeforeCreate(scope *gorm.Scope) error {

	err := scope.SetColumn("CreatedAt", time.Now())
	if err != nil {
		log.Fatalf("Error during obj creating: %v", err)
		return err
	}

	err = scope.SetColumn("UpdatedAt", time.Now())
	if err != nil {
		log.Fatalf("Error during obj creating: %v", err)
		return err
	}

	err = scope.SetColumn("ID", uuid.NewV4().String())
	if err != nil {
		log.Fatalf("Error during obj creating: %v", err)
		return err
	}

	return nil
}
