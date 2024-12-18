package cinema

import (
	"cinema-seating/internal/entity"
	"cinema-seating/pkg/errors"
	"context"
	"fmt"
	"net/http"
)

var (
	ErrSeatOutOfBounds = errors.NewError(
		http.StatusBadRequest,
		errors.CodeBadRequest,
		"seat out of bounds",
		"seat out of bounds",
	)
)

type UseCase struct {
	repo CinemaRepo
}

func New(repo CinemaRepo) *UseCase {
	return &UseCase{repo: repo}
}

func (uc *UseCase) ConfigureCinema(ctx context.Context, cinema *entity.Cinema) error {
	cinema.Seats = make([][]entity.Seat, cinema.Rows)
	for i := 0; i < cinema.Rows; i++ {
		cinema.Seats[i] = make([]entity.Seat, cinema.Columns)
		for j := 0; j < cinema.Columns; j++ {
			cinema.Seats[i][j] = entity.Seat{
				Row:    i,
				Column: j,
				Status: entity.Available,
			}
		}
	}

	err := uc.repo.SaveCinemaData(ctx, cinema)
	if err != nil {
		return err
	}

	return nil
}

func (uc *UseCase) GetAvailableSeats(ctx context.Context) ([]entity.AvailableSeat, error) {
	cinema, err := uc.repo.LoadCinemaData(ctx)
	if err != nil {
		return nil, err
	}

	availableSeats := make([]entity.AvailableSeat, 0)
	for i := 0; i < cinema.Rows; i++ {
		for j := 0; j < cinema.Columns; j++ {
			seatName := fmt.Sprintf("%c%d", 'A'+i, j+1) // A1, B1, ...
			availableSeats = append(availableSeats,
				entity.AvailableSeat{
					Name:   seatName,
					Row:    i,
					Col:    j,
					Status: cinema.Seats[i][j].Status,
				},
			)

		}
	}

	return availableSeats, nil
}

func (uc *UseCase) ReserveSeat(ctx context.Context, seat *entity.SeatAction) error {
	cinema, err := uc.repo.LoadCinemaData(ctx)
	if err != nil {
		return err
	}

	row, col := seat.Row, seat.Col
	if uc.isSeatOutOfBounds(row, col, cinema) {
		return ErrSeatOutOfBounds
	}

	if cinema.Seats[row][col].Status != entity.Available {
		return errors.NewError(
			http.StatusBadRequest,
			errors.CodeBadRequest,
			"seat is not available",
			"seat is not available",
		)
	}

	if !cinema.IsValidPlacement(row, row) {
		return errors.NewError(
			http.StatusBadRequest,
			errors.CodeBadRequest,
			"invalid placement",
			"invalid placement",
		)
	}

	cinema.Seats[row][col].Status = entity.Reserved
	err = uc.repo.SaveCinemaData(ctx, cinema)
	if err != nil {
		return err
	}

	return nil
}

func (uc *UseCase) CancelSeat(ctx context.Context, seat *entity.SeatAction) error {
	cinema, err := uc.repo.LoadCinemaData(ctx)
	if err != nil {
		return err
	}

	row, col := seat.Row, seat.Col
	if uc.isSeatOutOfBounds(row, col, cinema) {
		return ErrSeatOutOfBounds
	}

	if cinema.Seats[row][col].Status != entity.Reserved {
		return errors.NewError(
			http.StatusBadRequest,
			errors.CodeBadRequest,
			"seat is not reserved",
			"seat is not reserved",
		)
	}

	cinema.Seats[row][col].Status = entity.Available
	err = uc.repo.SaveCinemaData(ctx, cinema)
	if err != nil {
		return err
	}

	return nil
}

// isSeatOutOfBounds checks if the given row and column indices are out of the cinema's bounds.
func (uc *UseCase) isSeatOutOfBounds(row, col int, cinema *entity.Cinema) bool {
	return row < 0 || row >= cinema.Rows || col < 0 || col >= cinema.Columns
}
