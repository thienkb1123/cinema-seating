package repo

import (
	"cinema-seating/internal/entity"
	"context"
	"encoding/json"
	"os"
	"sync"
)

const (
	filePath = "data/cinema_data.json"
)

var (
	mu sync.Mutex
)

type CinemaRepo struct {
}

func NewJSON() *CinemaRepo {
	return &CinemaRepo{}
}

// LoadCinemaData đọc dữ liệu từ file JSON
func (r *CinemaRepo) LoadCinemaData(ctx context.Context) (*entity.Cinema, error) {
	mu.Lock()
	defer mu.Unlock()

	var cinema entity.Cinema

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return &cinema, nil
	}

	file, err := os.Open(filePath)
	if err != nil {
		return &cinema, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&cinema); err != nil {
		return &cinema, err
	}

	return &cinema, nil
}

// SaveCinemaData write data to file JSON
func (r *CinemaRepo) SaveCinemaData(ctx context.Context, cinema *entity.Cinema) error {
	mu.Lock()
	defer mu.Unlock()
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // Format data for read
	return encoder.Encode(cinema)
}
