package main

import (
	"fmt"
	"log"

	"github.com/elleven11/patient_api/api/models"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (srv *Server) Init(DbUser, DbPassword, DbPort, DbHost, DbName string) {
	var err error

	// DBURL is PostgreSQL compliant
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
	srv.DB, err = gorm.Open("postgres", DBURL)
	if err != nil {
		fmt.Println("Could not connect to DB")
		log.Fatal("Error: ", err)
	} else {
		fmt.Println("Successfully connected to DB!")
	}

	// migrate database
	srv.DB.Debug().AutoMigrate(&models.User{}, &models.Patient{})

	srv.Router = mux.NewRouter()
}
