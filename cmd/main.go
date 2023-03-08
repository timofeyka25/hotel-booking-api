package main

import (
	"database/sql"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"hotel-booking-app/dao"
	"log"
	"time"
)

func main() {
	connStr := "user=timofeyka.com.03 password=5jZF7HxWAXob dbname=hotelbookingdb host=ep-ancient-term-974725-pooler.eu-central-1.aws.neon.tech sslmode=verify-full"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	db.SetConnMaxLifetime(10 * time.Minute)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to db")
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
		log.Println("Disconnected from db")
	}(db)

	//user := domain.NewUser("Tymofii", "timofeyka.com.03@gmail.com", "123456", "6971c4dd-f5a3-4e3c-a90a-4feff2980aa9")
	//userDAO := dao.NewUserDAO(db)
	//err = userDAO.Create(user)
	//err = userDAO.Delete(uuid.MustParse("47385d1f-c63f-43b3-a3b5-a0104828a605"))
	hotelDAO := dao.NewHotelDAO(db)
	//err = hotelDAO.Create(domain.NewHotel("7 абщага", "Металістів 3", "dormitory"))
	h, err := hotelDAO.Read(uuid.MustParse("c3c16ec2-ccd8-45b9-a54c-be7aed603f7d"))
	if err != nil {
		log.Println(err)
	}
	log.Printf("%+v", *h)
}
