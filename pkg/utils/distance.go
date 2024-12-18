package utils

import "math"

// ManhattanDistance calculates the Manhattan distance between two points.
func ManhattanDistance(x1, y1, x2, y2 int) int {
	return int(math.Abs(float64(x1-x2)) + math.Abs(float64(y1-y2)))
}
