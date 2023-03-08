package dao

import (
	"database/sql"
	"github.com/google/uuid"
	"hotel-booking-app/domain"
)

type ReservationDAO interface {
	Create(domain.Reservation) error
	Read(uuid.UUID) (*domain.Reservation, error)
	Update(domain.Reservation) error
	Delete(uuid.UUID) error
}

type reservationDAO struct {
	db *sql.DB
}

func NewReservationDAO(db *sql.DB) *reservationDAO {
	return &reservationDAO{db: db}
}

func (dao reservationDAO) Create(r domain.Reservation) error {
	_, err := dao.db.Exec("INSERT INTO reservations (id, user_id, room_id, check_in_date, check_out_date, status)"+
		" VALUES ($1, $2, $3, $4, $5, $6)", r.Id, r.UserId, r.RoomId, r.CheckInDate, r.CheckOutDate, r.Status)

	return err
}

func (dao reservationDAO) Read(id uuid.UUID) (*domain.Reservation, error) {
	var r domain.Reservation

	err := dao.db.QueryRow("SELECT * FROM reservations WHERE id = $1", id).
		Scan(&r.Id, &r.UserId, &r.RoomId, &r.CheckInDate, &r.CheckOutDate, &r.Status, &r.PaymentStatus)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &r, nil
}

func (dao reservationDAO) Update(r domain.Reservation) error {
	_, err := dao.db.Exec("UPDATE reservations SET user_id = $2, room_id = $3, check_in_date = $4, check_out_date = $5, "+
		"status = $6, payment_status = $7 WHERE id = $1",
		r.Id, r.UserId, r.RoomId, r.CheckInDate, r.CheckOutDate, r.Status, r.PaymentStatus)

	return err
}

func (dao reservationDAO) Delete(id uuid.UUID) error {
	_, err := dao.db.Exec("DELETE FROM reservations WHERE id = $1", id)

	return err
}
