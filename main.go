package main

import (
	"github.com/br4tech/schedule-backoffice/internal/repositories"
	"github.com/jinzhu/gorm"
)

func main(){
	db, err := gorm.Open("")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userRepo := repositories.NewUserRepository(db)
}