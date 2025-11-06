package entities

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// Entity represents a basic game entity
type Entity struct {
	ID              string
	X, Y            float64
	Width, Height   float64
	VelX, VelY      float64
	Sprite          *Sprite
	Health          *Health
}

// EntityInterface defines the interface for entities
type EntityInterface interface {
	Update()
	Draw(screen *ebiten.Image)
	GetID() string
	GetX() float64
	GetY() float64
	GetWidth() float64
	GetHeight() float64
}

// Update updates the entity position
func (e *Entity) Update() {
	e.X += e.VelX
	e.Y += e.VelY
}

// Draw draws the entity
func (e *Entity) Draw(screen *ebiten.Image) {
	if e.Sprite != nil && e.Sprite.Image != nil {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(e.X, e.Y)
		screen.DrawImage(e.Sprite.Image, op)
	}
}

// GetID returns the entity ID
func (e *Entity) GetID() string {
	return e.ID
}

// GetX returns the X position
func (e *Entity) GetX() float64 {
	return e.X
}

// GetY returns the Y position
func (e *Entity) GetY() float64 {
	return e.Y
}

// GetWidth returns the width
func (e *Entity) GetWidth() float64 {
	return e.Width
}

// GetHeight returns the height
func (e *Entity) GetHeight() float64 {
	return e.Height
}

