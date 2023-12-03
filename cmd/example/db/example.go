package main

import (
	"github.com/rs/zerolog/log"
	"github.com/sourcehaven/mypass-api/internal/db"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {

	gormDb, err := gorm.Open(sqlite.Open("::memory:?cache=shared")) //, &gorm.Cfg{Logger: dbLogger}
	if err != nil {
		panic(err)
	}

	db := db.TransactDb{DB: gormDb}
	err = db.Init()
	if err != nil {
		panic(err)
	}

	user := &db.User{
		Username: "user",
		Password: []byte("password"),
	}
	userWithMeta, err := db.CreateUser(user)

	if err != nil {
		panic(err)
	}

	log.Print(userWithMeta)
	//db.GetUser()
	//db.GetUserById()
	//db.UpdateUser()
	//db.DeleteUser()
	//db.GetUserById()

}
