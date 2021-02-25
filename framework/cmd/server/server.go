package main

import (
	"fmt"
	"log"
	"github.com/service-user/application/repositories"
	"github.com/service-user/domain"
	"github.com/service-user/framework/utils"
)

func main() {

	db := utils.ConnectDB()

	user := domain.User{
		Name:     "Luiz Carlos Moitinho",
		Email:    "luizcarlos_costam@hotmail.com",
		Password: "123456",
	}

	userRepo := repositories.UserRepositoryDb{Db: db}
	result, err := userRepo.Insert(&user)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}
