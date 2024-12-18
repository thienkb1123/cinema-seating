package cinema

import (
	"cinema-seating/internal/entity"
	"context"
)

type Cinema interface {
	ConfigureCinema(ctx context.Context, cinema *entity.Cinema) error
	GetAvailableSeats(ctx context.Context) ([]entity.AvailableSeat, error)
	ReserveSeat(ctx context.Context, seat *entity.SeatAction) error
	CancelSeat(ctx context.Context, seat *entity.SeatAction) error
}

type CinemaRepo interface {
	LoadCinemaData(ctx context.Context) (*entity.Cinema, error)
	SaveCinemaData(ctx context.Context, cinema *entity.Cinema) error
}
