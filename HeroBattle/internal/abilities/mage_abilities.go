package abilities

import (
	"errors"
	"game/internal/entities"
)

// FireballAbility creates a projectile that explodes
type FireballAbility struct {
	BaseAbility
	Damage    int
	Range     float64
	AOERadius float64
	Speed     float64
}

// NewFireballAbility creates a new fireball ability
func NewFireballAbility() *FireballAbility {
	return &FireballAbility{
		BaseAbility: BaseAbility{
			Name:     "Fireball",
			Cooldown: 60,
			ManaCost: 20,
		},
		Damage:    25,
		Range:     300,
		AOERadius: 50,
		Speed:     400.0,
	}
}

// Execute executes the fireball ability
func (f *FireballAbility) Execute(hero *entities.Hero, g GameInterface) error {
	if !f.IsReady() {
		return errors.New("ability on cooldown")
	}
	if hero.Mana.Current < f.ManaCost {
		return errors.New("not enough mana")
	}
	
	hero.Mana.Current -= f.ManaCost
	f.CurrentCooldown = f.Cooldown
	
	targetX := hero.X + f.Range
	targetY := hero.Y
	
	enemies := g.GetEnemies()
	minDist := f.Range + 100
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
		f.Speed,
		f.Damage,
		hero,
	)
	projectile.AOERadius = f.AOERadius
	projectile.IsExplosive = true
	
	g.AddProjectile(projectile)
	return nil
}

// FrostNovaAbility creates an AOE that slows enemies
type FrostNovaAbility struct {
	BaseAbility
	Damage      int
	AOERadius   float64
	SlowPercent float64
}

// NewFrostNovaAbility creates a new frost nova ability
func NewFrostNovaAbility() *FrostNovaAbility {
	return &FrostNovaAbility{
		BaseAbility: BaseAbility{
			Name:     "Frost Nova",
			Cooldown: 180,
			ManaCost: 30,
		},
		Damage:      15,
		AOERadius:   100,
		SlowPercent: 0.5,
	}
}

// Execute executes the frost nova ability
func (f *FrostNovaAbility) Execute(hero *entities.Hero, g GameInterface) error {
	if !f.IsReady() {
		return errors.New("ability on cooldown")
	}
	if hero.Mana.Current < f.ManaCost {
		return errors.New("not enough mana")
	}
	
	hero.Mana.Current -= f.ManaCost
	f.CurrentCooldown = f.Cooldown
	
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
		
		if dist <= f.AOERadius {
			enemy.Health.Current -= f.Damage
			if enemy.Health.Current < 0 {
				enemy.Health.Current = 0
			}
		}
	}
	
	return nil
}

type ArcaneShieldAbility struct {
	BaseAbility
	Duration    int
	DamageReduction float64
}

// NewArcaneShieldAbility creates a new arcane shield ability
func NewArcaneShieldAbility() *ArcaneShieldAbility {
	return &ArcaneShieldAbility{
		BaseAbility: BaseAbility{
			Name:     "Arcane Shield",
			Cooldown: 300,
			ManaCost: 25,
		},
		Duration:        300,
		DamageReduction: 0.5,
	}
}

func (a *ArcaneShieldAbility) Execute(hero *entities.Hero, g GameInterface) error {
	if !a.IsReady() {
		return errors.New("ability on cooldown")
	}
	if hero.Mana.Current < a.ManaCost {
		return errors.New("not enough mana")
	}
	
	hero.Mana.Current -= a.ManaCost
	a.CurrentCooldown = a.Cooldown
	return nil
}

// ManaRegenPassive regenerates mana over time
type ManaRegenPassive struct {
	BaseAbility
	RegenRate float64 // mana per second
}

// NewManaRegenPassive creates a new mana regen passive
func NewManaRegenPassive() *ManaRegenPassive {
	return &ManaRegenPassive{
		BaseAbility: BaseAbility{
			Name: "Mana Regeneration",
		},
		RegenRate: 10.0, // 10 mana per second
	}
}

func (m *ManaRegenPassive) Execute(hero *entities.Hero, g GameInterface) error {
	return nil
}

func (m *ManaRegenPassive) Update(deltaTime float64) {
}

// MeteorStormAbility rains meteors over an area
type MeteorStormAbility struct {
	BaseAbility
	Damage       int
	AOERadius    float64
	MeteorCount  int
}

// NewMeteorStormAbility creates a new meteor storm ability
func NewMeteorStormAbility() *MeteorStormAbility {
	return &MeteorStormAbility{
		BaseAbility: BaseAbility{
			Name:     "Meteor Storm",
			Cooldown: 600, // 10 seconds
			ManaCost: 0,   // Ultimate uses energy, not mana
		},
		Damage:      40,
		AOERadius:   150,
		MeteorCount: 5,
	}
}

func (m *MeteorStormAbility) Execute(hero *entities.Hero, g GameInterface) error {
	return nil
}

