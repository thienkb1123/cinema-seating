package entity

import (
	"cinema-seating/pkg/utils"
)

// SeatStatus represents the status of a seat
// swagger:enum SeatStatus
type SeatStatus int

const (
	// Available indicates the seat is available
	Available SeatStatus = iota // 0
	// Reserved indicates the seat is reserved
	Reserved // 1
)

type Seat struct {
	Row    int        `json:"row"`
	Column int        `json:"column"`
	Status SeatStatus `json:"status"` // Status 0: Available, 1: Reserved
}

type Cinema struct {
	Rows        int      `json:"rows" validate:"required"`
	Columns     int      `json:"columns" validate:"required"`
	MinDistance int      `json:"minDistance" default:"0"`
	Seats       [][]Seat `json:"seats"`
}

// IsValidPlacement check if a seat can be placed at (row, col)
func (c *Cinema) IsValidPlacement(newRow, newCol int) bool {
	for _, row := range c.Seats {
		for _, seat := range row {
			if seat.Status == Reserved {
				distance := utils.ManhattanDistance(seat.Row, seat.Column, newRow, newCol)
				if distance <= c.MinDistance {
					return false
				}
			}
		}
	}
	return true
}

type AvailableSeat struct {
	Name   string     `json:"name"`
	Row    int        `json:"row"`
	Col    int        `json:"col"`
	Status SeatStatus `json:"status"` // Status 0: Available, 1: Reserved
}

type SeatAction struct {
	Row int `json:"row" default:"0"` // Row is required but can be 0
	Col int `json:"col" default:"0"` // Col is required but can be 0
}
