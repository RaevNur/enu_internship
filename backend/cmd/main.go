package main

import (
	"fmt"
	"log"
	"os"

	"enu_internship/internal/database"
	"enu_internship/internal/repository"
	"enu_internship/internal/service"
	"enu_internship/internal/web"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("configs/.env")
	if err != nil {
		log.Fatalf("unable to load env file %v ", err)
	}

	dbUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&timeout=1s",
		os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DB"))

	db, err := database.InitDB(dbUrl)
	if err != nil {
		log.Fatalf("unable to connect database %v ", err)
	}

	// create repos and services
	repos := repository.NewRepo(db)
	services := service.NewService(repos)

	// create handlers
	handlers, err := web.NewMainHandler(services)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	server := new(web.Server)
	if err := server.Run(os.Getenv("SERVER_PORT"), handlers.InitRoutes()); err != nil {
		log.Printf("ERROR: %s\n", err)
		os.Exit(1)
	}
	os.Exit(0)
}
