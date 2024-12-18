package cinema

import (
	"cinema-seating/internal/entity"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockCinemaRepo is a mock implementation of the CinemaRepo interface.
type MockCinemaRepo struct {
	mock.Mock
}

func (m *MockCinemaRepo) SaveCinemaData(ctx context.Context, cinema *entity.Cinema) error {
	args := m.Called(ctx, cinema)
	return args.Error(0)
}

func (m *MockCinemaRepo) LoadCinemaData(ctx context.Context) (*entity.Cinema, error) {
	args := m.Called(ctx)
	return args.Get(0).(*entity.Cinema), args.Error(1)
}

func TestConfigureCinema(t *testing.T) {
	mockRepo := new(MockCinemaRepo)
	uc := New(mockRepo)

	cinema := &entity.Cinema{
		Rows:    5,
		Columns: 5,
	}

	mockRepo.On("SaveCinemaData", mock.Anything, cinema).Return(nil)

	err := uc.ConfigureCinema(context.Background(), cinema)
	assert.NoError(t, err)
	assert.Equal(t, 5, len(cinema.Seats))
	assert.Equal(t, 5, len(cinema.Seats[0]))
	mockRepo.AssertExpectations(t)
}

func TestGetAvailableSeats(t *testing.T) {
	mockRepo := new(MockCinemaRepo)
	uc := New(mockRepo)

	cinema := &entity.Cinema{
		Rows:    5,
		Columns: 5,
		Seats:   make([][]entity.Seat, 5),
	}
	for i := 0; i < 5; i++ {
		cinema.Seats[i] = make([]entity.Seat, 5)
	}

	mockRepo.On("LoadCinemaData", mock.Anything).Return(cinema, nil)

	seats, err := uc.GetAvailableSeats(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, 25, len(seats))
	mockRepo.AssertExpectations(t)
}

func TestReserveSeat(t *testing.T) {
	mockRepo := new(MockCinemaRepo)
	uc := New(mockRepo)

	cinema := &entity.Cinema{
		Rows:    5,
		Columns: 5,
		Seats:   make([][]entity.Seat, 5),
	}
	for i := 0; i < 5; i++ {
		cinema.Seats[i] = make([]entity.Seat, 5)
	}

	mockRepo.On("LoadCinemaData", mock.Anything).Return(cinema, nil)
	mockRepo.On("SaveCinemaData", mock.Anything, cinema).Return(nil)

	seatAction := &entity.SeatAction{Row: 1, Col: 1}
	err := uc.ReserveSeat(context.Background(), seatAction)
	assert.NoError(t, err)
	assert.Equal(t, entity.Reserved, cinema.Seats[1][1].Status)
	mockRepo.AssertExpectations(t)
}

func TestCancelSeat(t *testing.T) {
	mockRepo := new(MockCinemaRepo)
	uc := New(mockRepo)

	cinema := &entity.Cinema{
		Rows:    5,
		Columns: 5,
		Seats:   make([][]entity.Seat, 5),
	}
	for i := 0; i < 5; i++ {
		cinema.Seats[i] = make([]entity.Seat, 5)
	}
	cinema.Seats[1][1].Status = entity.Reserved

	mockRepo.On("LoadCinemaData", mock.Anything).Return(cinema, nil)
	mockRepo.On("SaveCinemaData", mock.Anything, cinema).Return(nil)

	seatAction := &entity.SeatAction{Row: 1, Col: 1}
	err := uc.CancelSeat(context.Background(), seatAction)
	assert.NoError(t, err)
	assert.Equal(t, entity.Available, cinema.Seats[1][1].Status)
	mockRepo.AssertExpectations(t)
}
