package usecase

import (
	"context"
	"github.com/google/uuid"
	"hotel-booking-app/internal/dao"
	"hotel-booking-app/internal/domain"
	"hotel-booking-app/pkg/customErrors"
	"hotel-booking-app/pkg/db"
	"time"
)

type PaymentUseCase interface {
	PayForReservation(ctx context.Context, params CreatePaymentParams) (uuid.UUID, error)
}

type paymentUseCase struct {
	paymentDAO         dao.PaymentDAO
	reservationDAO     dao.ReservationDAO
	transactionManager *db.TransactionManager
}

func NewPaymentUseCase(
	paymentDAO dao.PaymentDAO,
	reservationDAO dao.ReservationDAO,
	transactionManager *db.TransactionManager,
) *paymentUseCase {
	return &paymentUseCase{
		paymentDAO:         paymentDAO,
		reservationDAO:     reservationDAO,
		transactionManager: transactionManager,
	}
}

func (uc paymentUseCase) PayForReservation(ctx context.Context, params CreatePaymentParams) (uuid.UUID, error) {
	reservation, err := uc.reservationDAO.GetById(ctx, params.ReservationId)
	if err != nil {
		return uuid.Nil, err
	}
	if reservation.Status != domain.PENDING {
		return uuid.Nil, customErrors.NewStatusError("Payment for this booking is not possible because wrong status")
	}
	if reservation.PaymentStatus == domain.PAID {
		return uuid.Nil, customErrors.NewStatusError("This reservation has already been paid for")
	}
	if reservation.CheckInDate.Before(time.Now()) {
		return uuid.Nil, customErrors.NewStatusError("Payment for this booking is not possible")
	}
	if reservation.Room.PricePerNight != params.Amount {
		return uuid.Nil, customErrors.NewStatusError("Wrong amount")
	}
	payment := domain.NewPayment(params.ReservationId, params.UserId, params.Amount)
	reservation.PaymentStatus = domain.PAID
	if err = uc.transactionManager.WithinTransaction(ctx, func(ctx context.Context) error {
		if err = uc.paymentDAO.Create(ctx, payment); err != nil {
			return err
		}
		if err = uc.reservationDAO.Update(ctx, reservation); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return uuid.Nil, err
	}
	return payment.Id, nil
}

type CreatePaymentParams struct {
	ReservationId uuid.UUID
	UserId        uuid.UUID
	Amount        float64
}
