package abilities

import (
	"errors"
	"game/internal/entities"
	"image/color"
)

// QuickShotAbility fires a fast arrow
type QuickShotAbility struct {
	BaseAbility
	Damage int
	Range  float64
	Speed  float64
}

// NewQuickShotAbility creates a new quick shot ability
func NewQuickShotAbility() *QuickShotAbility {
	return &QuickShotAbility{
		BaseAbility: BaseAbility{
			Name:     "Quick Shot",
			Cooldown: 20,
			ManaCost: 8,
		},
		Damage: 15,
		Range:  400,
		Speed:  500,
	}
}

// Execute executes the quick shot ability
func (q *QuickShotAbility) Execute(hero *entities.Hero, g GameInterface) error {
	if !q.IsReady() {
		return errors.New("ability on cooldown")
	}
	if hero.Mana.Current < q.ManaCost {
		return errors.New("not enough mana")
	}
	
	hero.Mana.Current -= q.ManaCost
	q.CurrentCooldown = q.Cooldown
	
	var targetX, targetY float64 = hero.X + q.Range, hero.Y
	enemies := g.GetEnemies()
	minDist := q.Range + 100
	
	for _, enemy := range enemies {
		if enemy.Health == nil || enemy.Health.Current <= 0 {
			continue
		}
		
		dist := entities.GetDistance(hero.X, hero.Y, enemy.X, enemy.Y)
		if dist < minDist {
			minDist = dist
			targetX = enemy.X + enemy.Width/2
			targetY = enemy.Y + enemy.Height/2
		}
	}
	
	projectile := entities.NewProjectile(
		hero.X, hero.Y,
		targetX, targetY,
		q.Speed,
		q.Damage,
		hero,
	)
	projectile.Sprite = entities.NewSpriteFromAsset("arrow")
	if projectile.Sprite == nil || projectile.Sprite.Image == nil {
		projectile.Sprite = entities.NewSpriteFromColor(int(projectile.Width), int(projectile.Height), color.RGBA{150, 75, 0, 255})
	}
	
	g.AddProjectile(projectile)
	return nil
}

// ArrowRainAbility creates an AOE rain of arrows
type ArrowRainAbility struct {
	BaseAbility
	Damage      int
	AOERadius   float64
	ArrowCount  int
}

// NewArrowRainAbility creates a new arrow rain ability
func NewArrowRainAbility() *ArrowRainAbility {
	return &ArrowRainAbility{
		BaseAbility: BaseAbility{
			Name:     "Arrow Rain",
			Cooldown: 240,
			ManaCost: 25,
		},
		Damage:     12,
		AOERadius:  120,
		ArrowCount: 8,
	}
}

// Execute executes the arrow rain ability
func (a *ArrowRainAbility) Execute(hero *entities.Hero, g GameInterface) error {
	if !a.IsReady() {
		return errors.New("ability on cooldown")
	}
	if hero.Mana.Current < a.ManaCost {
		return errors.New("not enough mana")
	}
	
	hero.Mana.Current -= a.ManaCost
	a.CurrentCooldown = a.Cooldown
	
	enemies := g.GetEnemies()
	heroCenterX := hero.X + hero.Width/2
	heroCenterY := hero.Y + hero.Height/2
	
	for _, enemy := range enemies {
		if enemy.Health == nil || enemy.Health.Current <= 0 {
			continue
		}
		
		enemyCenterX := enemy.X + enemy.Width/2
		enemyCenterY := enemy.Y + enemy.Height/2
		dist := entities.GetDistance(heroCenterX, heroCenterY, enemyCenterX, enemyCenterY)
		
		if dist <= a.AOERadius {
			projectile := entities.NewProjectile(
				hero.X, hero.Y,
				enemyCenterX, enemyCenterY,
				400.0, // Speed
				a.Damage,
				hero,
			)
			// Use arrow sprite asset
			projectile.Sprite = entities.NewSpriteFromAsset("arrow")
			if projectile.Sprite == nil || projectile.Sprite.Image == nil {
				projectile.Sprite = entities.NewSpriteFromColor(int(projectile.Width), int(projectile.Height), color.RGBA{150, 75, 0, 255})
			}
			g.AddProjectile(projectile)
		}
	}
	
	return nil
}

// EvasionAbility dodges the next attack
type EvasionAbility struct {
	BaseAbility
	Duration int // in frames
}

// NewEvasionAbility creates a new evasion ability
func NewEvasionAbility() *EvasionAbility {
	return &EvasionAbility{
		BaseAbility: BaseAbility{
			Name:     "Evasion",
			Cooldown: 600,
			ManaCost: 20,
		},
		Duration: 600,
	}
}

func (e *EvasionAbility) Execute(hero *entities.Hero, g GameInterface) error {
	if !e.IsReady() {
		return errors.New("ability on cooldown")
	}
	if hero.Mana.Current < e.ManaCost {
		return errors.New("not enough mana")
	}
	
	hero.Mana.Current -= e.ManaCost
	e.CurrentCooldown = e.Cooldown
	return nil
}

// CriticalStrikePassive gives chance for double damage
type CriticalStrikePassive struct {
	BaseAbility
	CritChance   float64
	CritMultiplier float64
}

// NewCriticalStrikePassive creates a new critical strike passive
func NewCriticalStrikePassive() *CriticalStrikePassive {
	return &CriticalStrikePassive{
		BaseAbility: BaseAbility{
			Name: "Critical Strike",
		},
		CritChance:    0.3,
		CritMultiplier: 2.0,
	}
}

func (c *CriticalStrikePassive) Execute(hero *entities.Hero, g GameInterface) error {
	return nil
}

// PerfectShotAbility guarantees crit with high damage
type PerfectShotAbility struct {
	BaseAbility
	Damage      int
	Range       float64
	Speed       float64
	DamageMultiplier float64
}

// NewPerfectShotAbility creates a new perfect shot ability
func NewPerfectShotAbility() *PerfectShotAbility {
	return &PerfectShotAbility{
		BaseAbility: BaseAbility{
			Name:     "Perfect Shot",
			Cooldown: 600,
			ManaCost: 0,
		},
		Damage:          50,
		Range:           500,
		Speed:           600,
		DamageMultiplier: 3.0,
	}
}

func (p *PerfectShotAbility) Execute(hero *entities.Hero, g GameInterface) error {
	return nil
}

