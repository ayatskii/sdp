package entities

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

// Enemy represents an enemy entity
type Enemy struct {
	Entity
	Type        string
	AttackRange float64
	ViewRange   float64
	Attack      int
	Defense     int
	Speed       float64
	AttackCooldown int
	CurrentAttackCooldown int
}

// NewEnemy creates a new enemy
func NewEnemy(enemyType string, x, y float64) *Enemy {
	e := &Enemy{
		Entity: Entity{
			ID:     enemyType,
			X:      x,
			Y:      y,
			Width:  32,
			Height: 32,
		},
		Type:        enemyType,
		AttackRange: 50,
		ViewRange:   300,
		Attack:      5,
		Defense:     2,
		Speed:       100.0 / 60.0,
		AttackCooldown: 120,
		CurrentAttackCooldown: 0,
	}
	
	switch enemyType {
	case "Melee":
		e.Health = &Health{Current: 50, Max: 50}
		e.Attack = 8
		e.Defense = 3
		e.Sprite = NewSpriteFromAsset("melee")
		if e.Sprite == nil || e.Sprite.Image == nil {
			e.SetSprite(color.RGBA{200, 50, 50, 255})
		}
	case "Ranged":
		e.Health = &Health{Current: 30, Max: 30}
		e.Attack = 6
		e.Defense = 1
		e.AttackRange = 200
		e.Sprite = NewSpriteFromAsset("ranged")
		if e.Sprite == nil || e.Sprite.Image == nil {
			e.SetSprite(color.RGBA{50, 200, 50, 255})
		}
	case "Tank":
		e.Health = &Health{Current: 100, Max: 100}
		e.Attack = 10
		e.Defense = 8
		e.Speed = 50.0 / 60.0 // Slower
		e.Sprite = NewSpriteFromAsset("tank")
		if e.Sprite == nil || e.Sprite.Image == nil {
			e.SetSprite(color.RGBA{50, 50, 200, 255})
		}
	}
	
	return e
}

// Update updates the enemy
func (e *Enemy) Update() {
	e.Entity.Update()
	if e.CurrentAttackCooldown > 0 {
		e.CurrentAttackCooldown--
	}
}

// Draw draws the enemy
func (e *Enemy) Draw(screen *ebiten.Image) {
	e.Entity.Draw(screen)
}

// SetSprite sets the enemy's sprite
func (e *Enemy) SetSprite(c color.Color) {
	e.Sprite = NewSpriteFromColor(int(e.Width), int(e.Height), c)
}

