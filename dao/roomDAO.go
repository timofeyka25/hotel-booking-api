package dao

import (
	"database/sql"
	"github.com/google/uuid"
	"hotel-booking-app/domain"
)

type RoomDAO interface {
	Create(domain.Room) error
	Read(uuid.UUID) (*domain.Room, error)
	Update(domain.Room) error
	Delete(uuid.UUID) error
}

type roomDAO struct {
	db *sql.DB
}

func NewRoomDAO(db *sql.DB) *roomDAO {
	return &roomDAO{db: db}
}

func (dao roomDAO) Create(room domain.Room) error {
	_, err := dao.db.Exec("INSERT INTO rooms (id, hotel_id, room_type, max_occupancy, price_per_night) VALUES ($1, $2, $3, $4, $5)",
		room.Id, room.HotelId, room.RoomType, room.MaxOccupancy, room.PricePerNight)

	return err
}

func (dao roomDAO) Read(id uuid.UUID) (*domain.Room, error) {
	row := dao.db.QueryRow("SELECT id, hotel_id, room_type, max_occupancy, price_per_night FROM rooms WHERE id = $1", id)

	var room domain.Room
	err := row.Scan(&room.Id, &room.HotelId, &room.RoomType, &room.MaxOccupancy, &room.PricePerNight)
	if err == sql.ErrNoRows {
		return nil, nil // no room found
	} else if err != nil {
		return nil, err
	}

	return &room, nil
}

func (dao roomDAO) Update(room domain.Room) error {
	_, err := dao.db.Exec("UPDATE rooms SET hotel_id = $1, room_type = $2, max_occupancy = $3, price_per_night = $4 WHERE id = $5",
		room.HotelId, room.RoomType, room.MaxOccupancy, room.PricePerNight, room.Id)

	return err
}

func (dao roomDAO) Delete(id uuid.UUID) error {
	_, err := dao.db.Exec("DELETE FROM rooms WHERE id = $1", id)

	return err
}
