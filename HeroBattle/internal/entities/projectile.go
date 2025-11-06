package entities

import (
	"image/color"
	"math"
)

// Projectile represents a projectile in the game
type Projectile struct {
	Entity
	TargetX    float64
	TargetY    float64
	Speed      float64
	Damage     int
	AOERadius  float64
	Owner      *Hero
	Lifetime   int
	MaxLifetime int
	IsExplosive bool
}

// NewProjectile creates a new projectile
func NewProjectile(x, y, targetX, targetY, speed float64, damage int, owner *Hero) *Projectile {
	dx := targetX - x
	dy := targetY - y
	distance := math.Sqrt(dx*dx + dy*dy)
	
	velX := (dx / distance) * speed
	velY := (dy / distance) * speed
	
	p := &Projectile{
		Entity: Entity{
			ID:     "projectile",
			X:      x,
			Y:      y,
			Width:  8,
			Height: 8,
			VelX:   velX,
			VelY:   velY,
		},
		TargetX:    targetX,
		TargetY:    targetY,
		Speed:      speed,
		Damage:     damage,
		AOERadius:  0,
		Owner:      owner,
		Lifetime:   0,
		MaxLifetime: 300,
		IsExplosive: false,
	}
	
	p.Sprite = NewSpriteFromAsset("fireball")
	if p.Sprite == nil || p.Sprite.Image == nil {
		p.Sprite = NewSpriteFromColor(int(p.Width), int(p.Height), color.RGBA{255, 255, 100, 255})
	}
	return p
}

func (p *Projectile) Update() {
	p.Entity.Update()
	p.Lifetime++
	
	if p.Lifetime >= p.MaxLifetime {
		return
	}
	
	dx := p.TargetX - p.X
	dy := p.TargetY - p.Y
	distance := math.Sqrt(dx*dx + dy*dy)
	if distance < 10 {
		p.Lifetime = p.MaxLifetime
	}
}

func (p *Projectile) IsExpired() bool {
	return p.Lifetime >= p.MaxLifetime
}

