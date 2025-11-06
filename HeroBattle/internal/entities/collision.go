package entities

import "math"

// CheckRectCollision checks if two rectangles collide
func CheckRectCollision(a, b *Entity) bool {
	return a.X < b.X+b.Width &&
		a.X+a.Width > b.X &&
		a.Y < b.Y+b.Height &&
		a.Y+a.Height > b.Y
}

// CheckCircleCollision checks if two circles collide
func CheckCircleCollision(x1, y1, r1, x2, y2, r2 float64) bool {
	dx := x2 - x1
	dy := y2 - y1
	distance := math.Sqrt(dx*dx + dy*dy)
	return distance < (r1 + r2)
}

// GetDistance calculates distance between two points
func GetDistance(x1, y1, x2, y2 float64) float64 {
	dx := x2 - x1
	dy := y2 - y1
	return math.Sqrt(dx*dx + dy*dy)
}

