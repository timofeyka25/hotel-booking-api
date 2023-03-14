package handler

import (
	"github.com/google/uuid"
	"hotel-booking-app/internal/domain"
	"hotel-booking-app/internal/handler/dto"
	"hotel-booking-app/internal/usecase"
)

func toSignUpParams(dto *dto.SignUpRequestDTO) usecase.SignUpParams {
	return usecase.SignUpParams{
		Name:     dto.Name,
		Email:    dto.Email,
		Password: dto.Password,
	}
}

func toSignInParams(dto *dto.SignInRequestDTO) usecase.SignInParams {
	return usecase.SignInParams{
		Email:    dto.Email,
		Password: dto.Password,
	}
}

func toAddHotelParams(dto *dto.AddHotelDTO) usecase.AddHotelParams {
	return usecase.AddHotelParams{
		Name:        dto.Name,
		Location:    dto.Location,
		Description: dto.Description,
	}
}

func toAddRoomParams(id uuid.UUID, dto *dto.AddRoomDTO) usecase.AddRoomParams {
	return usecase.AddRoomParams{
		HotelId:       id,
		RoomType:      dto.RoomType,
		MaxOccupancy:  dto.MaxOccupancy,
		PricePerNight: dto.PricePerNight,
	}
}

func toUpdateHotelParams(id uuid.UUID, dto *dto.UpdateHotelDTO) usecase.UpdateHotelParams {
	return usecase.UpdateHotelParams{
		Id:          id,
		Name:        dto.Name,
		Location:    dto.Location,
		Description: dto.Description,
	}
}

func toUpdateRoomParams(id uuid.UUID, dto *dto.UpdateRoomDTO) usecase.UpdateRoomParams {
	return usecase.UpdateRoomParams{
		Id:            id,
		RoomType:      dto.RoomType,
		MaxOccupancy:  dto.MaxOccupancy,
		PricePerNight: dto.PricePerNight,
	}
}

func toCreateReservationParams(
	roomId, userId uuid.UUID,
	dto *dto.CreateReservationParsedDTO,
) usecase.CreateReservationParams {
	return usecase.CreateReservationParams{
		UserId:       userId,
		RoomId:       roomId,
		CheckInDate:  dto.CheckInDate,
		CheckOutDate: dto.CheckOutDate,
	}
}

func toUpdateReservationStatusParams(
	id uuid.UUID,
	dto *dto.UpdateReservationStatusDTO,
) usecase.UpdateReservationStatusParams {
	return usecase.UpdateReservationStatusParams{
		Id:     id,
		Status: dto.Status,
	}
}

func toCreatePaymentParams(reservationId, userId uuid.UUID, amount float64) usecase.CreatePaymentParams {
	return usecase.CreatePaymentParams{ReservationId: reservationId, UserId: userId, Amount: amount}
}

func toUpdateUserStatusParams(userId uuid.UUID, isActive bool) usecase.UpdateUserActiveParams {
	return usecase.UpdateUserActiveParams{UserId: userId, IsActive: isActive}
}

func toUpdateRoleParams(userId uuid.UUID, roleDto *dto.UpdateRoleDTO) usecase.UpdateUserRoleParams {
	return usecase.UpdateUserRoleParams{UserId: userId, Role: roleDto.Role}
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

func mapDtoRoom(room *domain.Room) *dto.RoomDTO {
	return &dto.RoomDTO{
		Id:            room.Id,
		Hotel:         mapDtoHotel(room.Hotel),
		RoomType:      room.RoomType,
		MaxOccupancy:  room.MaxOccupancy,
		PricePerNight: room.PricePerNight,
	}
}

func mapDtoRooms(rooms []*domain.Room) []*dto.RoomDTO {
	var dtoRooms []*dto.RoomDTO
	for _, room := range rooms {
		dtoRooms = append(dtoRooms, mapDtoRoom(room))
	}
	return dtoRooms
}

func mapDtoReservation(reservation *domain.Reservation) *dto.ReservationDTO {
	return &dto.ReservationDTO{
		Id:            reservation.Id,
		UserId:        reservation.UserId,
		Room:          mapDtoRoom(reservation.Room),
		CheckInDate:   reservation.CheckInDate,
		CheckOutDate:  reservation.CheckOutDate,
		Status:        reservation.Status,
		PaymentStatus: reservation.PaymentStatus,
	}
}

func mapDtoReservations(reservations []*domain.Reservation) []*dto.ReservationDTO {
	var dtoReservations []*dto.ReservationDTO
	for _, reservation := range reservations {
		dtoReservations = append(dtoReservations, mapDtoReservation(reservation))
	}
	return dtoReservations
}

func mapDtoPayment(payment *domain.Payment) *dto.PaymentDTO {
	return &dto.PaymentDTO{
		Id:            payment.Id,
		ReservationId: payment.ReservationId,
		UserId:        payment.UserId,
		Amount:        payment.Amount,
		PaymentTime:   payment.PaymentTime,
	}
}

func mapDtoPayments(payments []*domain.Payment) []*dto.PaymentDTO {
	var dtoPayments []*dto.PaymentDTO
	for _, payment := range payments {
		dtoPayments = append(dtoPayments, mapDtoPayment(payment))
	}
	return dtoPayments
}

func mapDtoUser(user *domain.User) *dto.UserDTO {
	return &dto.UserDTO{
		Id:       user.Id,
		Name:     user.Name,
		Email:    user.Email,
		IsActive: user.IsActive,
		Role: &dto.RoleDTO{
			Id:   user.Role.Id,
			Name: user.Role.Name,
		},
	}
}

func mapDtoUsers(users []*domain.User) []*dto.UserDTO {
	var dtoUsers []*dto.UserDTO
	for _, user := range users {
		dtoUsers = append(dtoUsers, mapDtoUser(user))
	}
	return dtoUsers
}
