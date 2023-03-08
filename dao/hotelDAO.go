package dao

import (
	"database/sql"
	"github.com/google/uuid"
	"hotel-booking-app/domain"
)

type HotelDAO interface {
	Create(domain.Hotel) error
	Read(uuid.UUID) (*domain.Hotel, error)
	Update(domain.Hotel) error
	Delete(uuid.UUID) error
}

type hotelDAO struct {
	db *sql.DB
}

func NewHotelDAO(db *sql.DB) *hotelDAO {
	return &hotelDAO{db: db}
}

func (dao hotelDAO) Create(hotel domain.Hotel) error {
	_, err := dao.db.Exec("INSERT INTO hotels (id, name, location, description) VALUES ($1, $2, $3, $4)",
		hotel.Id, hotel.Name, hotel.Location, hotel.Description)

	return err
}

func (dao hotelDAO) Read(id uuid.UUID) (*domain.Hotel, error) {
	row := dao.db.QueryRow("SELECT id, name, location, description FROM hotels WHERE id = $1", id)

	var hotel domain.Hotel
	err := row.Scan(&hotel.Id, &hotel.Name, &hotel.Location, &hotel.Description)
	if err == sql.ErrNoRows {
		return nil, nil // no hotel found
	} else if err != nil {
		return nil, err
	}

	return &hotel, nil
}

func (dao hotelDAO) Update(hotel domain.Hotel) error {
	_, err := dao.db.Exec("UPDATE hotels SET name = $2, location = $3, description = $4 WHERE id = $1",
		hotel.Id, hotel.Name, hotel.Location, hotel.Description)

	return err
}

func (dao hotelDAO) Delete(id uuid.UUID) error {
	_, err := dao.db.Exec("DELETE FROM hotels WHERE id = $1", id)

	return err
}
