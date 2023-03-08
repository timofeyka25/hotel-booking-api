package main

import (
	"hotel-booking-app/dao"
	"hotel-booking-app/handler"
	"hotel-booking-app/pkg/db"
	"hotel-booking-app/pkg/server"
	"hotel-booking-app/usecase"
	"log"
)

func main() {
	dbInstance, err := db.NewDB(db.Config{
		Username: "timofeyka.com.03",
		Password: "5jZF7HxWAXob",
		DbName:   "hotelbookingdb",
		Host:     "ep-ancient-term-974725-pooler.eu-central-1.aws.neon.tech",
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to db")

	app := server.NewHTTPServer(server.Config{ReadTimeout: "10"})
	userHandler := handler.NewUserHandler(usecase.NewUserUseCase(dao.NewUserDAO(dbInstance)))
	handlers := handler.NewHandler(userHandler)
	handlers.InitRoutes(app)
	err = app.Listen(":8000")
	if err != nil {
		log.Fatal(err)
	}
	err = dbInstance.Close()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Disconnected from db")
}
