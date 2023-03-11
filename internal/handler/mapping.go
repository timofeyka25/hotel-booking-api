package handler

import (
	"github.com/google/uuid"
	"hotel-booking-app/internal/domain"
	"hotel-booking-app/internal/handler/dto"
	"hotel-booking-app/internal/usecase"
)

func toSignUpParams(dto *dto.SignUpRequestDTO) usecase.SignUpParams {
	return usecase.SignUpParams{Name: dto.Name, Email: dto.Email, Password: dto.Password}
}

func toSignInParams(dto *dto.SignInRequestDTO) usecase.SignInParams {
	return usecase.SignInParams{Email: dto.Email, Password: dto.Password}
}

func toAddHotelParams(dto *dto.AddHotelReqDTO) usecase.AddHotelParams {
	return usecase.AddHotelParams{Name: dto.Name, Location: dto.Location, Description: dto.Description}
}

func toUpdateHotelParams(id uuid.UUID, dto *dto.UpdateHotelDTO) usecase.UpdateHotelParams {
	return usecase.UpdateHotelParams{Id: id, Name: dto.Name, Location: dto.Location, Description: dto.Description}
}

func mapDtoHotel(hotel *domain.Hotel) *dto.HotelDTO {
	return &dto.HotelDTO{
		Id:          hotel.Id,
		Name:        hotel.Name,
		Location:    hotel.Location,
		Description: hotel.Description,
	}
}

func mapDtoHotels(hotels []*domain.Hotel) []*dto.HotelDTO {
	var dtoHotels []*dto.HotelDTO
	for _, hotel := range hotels {
		dtoHotels = append(dtoHotels, mapDtoHotel(hotel))
	}
	return dtoHotels
}
