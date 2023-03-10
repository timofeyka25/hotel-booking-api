package domain

import "github.com/google/uuid"

type Hotel struct {
	Id          uuid.UUID `bun:",pk"`
	Name        string
	Location    string
	Description string
}

func NewHotel(name, location, description string) Hotel {
	return Hotel{
		Id:          uuid.New(),
		Name:        name,
		Location:    location,
		Description: description,
	}
}
