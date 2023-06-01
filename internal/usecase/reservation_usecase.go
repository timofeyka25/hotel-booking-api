package usecase

import (
	"context"
	"github.com/google/uuid"
	"hotel-booking-app/internal/dao"
	"hotel-booking-app/internal/domain"
	"hotel-booking-app/pkg/custom_errors"
	"time"
)

type ReservationUseCase interface {
	CreateReservation(ctx context.Context, params CreateReservationParams) (uuid.UUID, error)
	GetAllUserReservations(ctx context.Context, userId uuid.UUID) ([]*domain.Reservation, error)
	GetAllReservations(ctx context.Context) ([]*domain.Reservation, error)
	CancelUserReservation(ctx context.Context, reservationId, userId uuid.UUID) error
	UpdateStatus(ctx context.Context, params UpdateReservationStatusParams) error
}

type reservationUseCase struct {
	reservationDAO dao.ReservationDAO
}

func NewReservationUseCase(reservationDAO dao.ReservationDAO) *reservationUseCase {
	return &reservationUseCase{reservationDAO: reservationDAO}
}

func (uc reservationUseCase) CreateReservation(ctx context.Context, params CreateReservationParams) (uuid.UUID, error) {
	reservations, err := uc.reservationDAO.GetByRoomAndUserId(ctx, params.UserId, params.RoomId)
	if err != nil {
		return uuid.Nil, err
	}
	if reservations != nil {
		for _, reservation := range reservations {
			if reservation.CheckOutDate.After(time.Now()) && reservation.Status != domain.CANCELLED {
				return uuid.Nil, custom_errors.NewAlreadyReservedError("This room is already reserved by you")
			}
		}
	}
	reservation := domain.NewReservation(params.UserId,
		params.RoomId,
		params.CheckInDate,
		params.CheckOutDate,
		domain.PENDING,
		domain.UNPAID)
	if err = uc.reservationDAO.Create(ctx, reservation); err != nil {
		return uuid.Nil, err
	}
	return reservation.Id, nil
}

func (uc reservationUseCase) GetAllUserReservations(ctx context.Context, id uuid.UUID) ([]*domain.Reservation, error) {
	return uc.reservationDAO.GetByUserId(ctx, id)
}

func (uc reservationUseCase) GetAllReservations(ctx context.Context) ([]*domain.Reservation, error) {
	return uc.reservationDAO.GetAll(ctx)
}

func (uc reservationUseCase) CancelUserReservation(ctx context.Context, reservationId, userId uuid.UUID) error {
	reservation, err := uc.reservationDAO.GetById(ctx, reservationId)
	if err != nil {
		return err
	}
	if reservation.UserId != userId {
		return custom_errors.NewNotFoundError("You do not have this reservation")
	}
	if reservation.Status == domain.CANCELLED || reservation.Status == domain.COMPLETED {
		return custom_errors.NewStatusError("You cannot cancel this reservation")
	}
	reservation.Status = domain.CANCELLED
	return uc.reservationDAO.Update(ctx, reservation)
}

func (uc reservationUseCase) UpdateStatus(ctx context.Context, params UpdateReservationStatusParams) error {
	reservation, err := uc.reservationDAO.GetById(ctx, params.Id)
	if err != nil {
		return err
	}
	if reservation.Status == params.Status {
		return custom_errors.NewStatusError("Status matches current reservation status")
	}
	switch params.Status {
	case domain.CONFIRMED:
		if reservation.PaymentStatus != domain.PAID {
			return custom_errors.NewStatusError("The reservation must be paid for")
		}
		if reservation.Status != domain.PENDING {
			return custom_errors.NewStatusError("The reservation status is no longer PENDING.")
		}
	case domain.COMPLETED, domain.UNFULFILLED:
		if reservation.CheckOutDate.After(time.Now()) {
			return custom_errors.NewStatusError("The reservation has not yet been completed")
		}
		if reservation.Status != domain.CONFIRMED {
			return custom_errors.NewStatusError("First you need to confirm this reservation")
		}
	case domain.CANCELLED:
		if reservation.Status != domain.PENDING {
			return custom_errors.NewStatusError("The reservation status is no longer PENDING.")
		}
	default:
		if reservation.Status == domain.CANCELLED {
			return custom_errors.NewStatusError("The reservation already cancelled")
		}
		return custom_errors.NewStatusError("Wrong status")
	}
	reservation.Status = params.Status
	return uc.reservationDAO.Update(ctx, reservation)
}

type CreateReservationParams struct {
	UserId       uuid.UUID
	RoomId       uuid.UUID
	CheckInDate  time.Time
	CheckOutDate time.Time
}

type UpdateReservationStatusParams struct {
	Id     uuid.UUID
	Status string
}
