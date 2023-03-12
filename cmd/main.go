package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"hotel-booking-app/internal/dao"
	"hotel-booking-app/internal/handler"
	"hotel-booking-app/internal/usecase"
	database "hotel-booking-app/pkg/db"
	"hotel-booking-app/pkg/jwt"
	"hotel-booking-app/pkg/server"
	"hotel-booking-app/pkg/validator"
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"
)

//	@title		Backend API
//  @version 1.0
//	@host		localhost:8000
//	@BasePath	/

func main() {
	// load env variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// init db
	db, err := database.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	defer closeDB(db)

	// init helpers
	jwtGenerator := jwt.NewTokenGenerator(jwt.Config{SecretKey: os.Getenv("JWT_SECRET_KEY")})
	jwtValidator := jwt.NewTokenValidator(jwt.Config{SecretKey: os.Getenv("JWT_SECRET_KEY")})
	validate := validator.NewValidator()

	// init dao
	userDao := dao.NewUserDAO(db)
	roleDAO := dao.NewRoleDAO(db)
	hotelDAO := dao.NewHotelDAO(db)
	roomDAO := dao.NewRoomDAO(db)

	// init use cases
	userUseCase := usecase.NewUserUseCase(userDao, roleDAO, jwtGenerator)
	hotelUseCase := usecase.NewHotelUseCase(hotelDAO)
	roomUseCase := usecase.NewRoomUseCase(roomDAO)

	// init handlers
	userHandler := handler.NewUserHandler(userUseCase, validate)
	hotelHandler := handler.NewHotelHandler(hotelUseCase, validate)
	roomHandler := handler.NewRoomHandler(roomUseCase, validate)
	handlers := handler.NewHandler(userHandler, hotelHandler, roomHandler)

	// init app
	app := server.NewHTTPServer(jwtValidator)
	handlers.InitRoutes(app)

	// run app
	go func() {
		err = app.Listen(fmt.Sprintf("%s:%s", os.Getenv("SERVER_HOST"), os.Getenv("SERVER_PORT")))
		if err != nil {
			log.Println("Server unexpectedly stopped")
		}
	}()

	// shutdown app and db
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	signal.Notify(stop, syscall.SIGTERM)
	<-stop

	if err = app.Shutdown(); err != nil {
		log.Printf("error shutting down server %s", err)
	} else {
		log.Println("Server gracefully stopped")
	}
}

func closeDB(db io.Closer) {
	if err := db.Close(); err != nil {
		log.Println(errors.Wrap(err, "err closing db connection"))
	} else {
		log.Println("db connection gracefully closed")
	}
}
