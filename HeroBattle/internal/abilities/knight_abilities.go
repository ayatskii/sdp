package abilities

import (
	"errors"
	"game/internal/entities"
)

// SlashAbility performs a melee attack
type SlashAbility struct {
	BaseAbility
	Damage      int
	Range       float64
}

// NewSlashAbility creates a new slash ability
func NewSlashAbility() *SlashAbility {
	return &SlashAbility{
		BaseAbility: BaseAbility{
			Name:     "Slash",
			Cooldown: 30,
			ManaCost: 10,
		},
		Damage: 20,
		Range:  50,
	}
}

// Execute executes the slash ability
func (s *SlashAbility) Execute(hero *entities.Hero, g GameInterface) error {
	if !s.IsReady() {
		return errors.New("ability on cooldown")
	}
	if hero.Mana.Current < s.ManaCost {
		return errors.New("not enough mana")
	}
	
	hero.Mana.Current -= s.ManaCost
	s.CurrentCooldown = s.Cooldown
	
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
		
		if dist <= s.Range {
			enemy.Health.Current -= s.Damage
			if enemy.Health.Current < 0 {
				enemy.Health.Current = 0
			}
		}
	}
	
	return nil
}

// ShieldBashAbility stuns and damages enemies
type ShieldBashAbility struct {
	BaseAbility
	Damage      int
	Range       float64
	StunDuration int // in frames
}

// NewShieldBashAbility creates a new shield bash ability
func NewShieldBashAbility() *ShieldBashAbility {
	return &ShieldBashAbility{
		BaseAbility: BaseAbility{
			Name:     "Shield Bash",
			Cooldown: 120,
			ManaCost: 15,
		},
		Damage:       25,
		Range:        60,
		StunDuration: 60,
	}
}

// Execute executes the shield bash ability
func (s *ShieldBashAbility) Execute(hero *entities.Hero, g GameInterface) error {
	if !s.IsReady() {
		return errors.New("ability on cooldown")
	}
	if hero.Mana.Current < s.ManaCost {
		return errors.New("not enough mana")
	}
	
	hero.Mana.Current -= s.ManaCost
	s.CurrentCooldown = s.Cooldown
	
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
		
		if dist <= s.Range {
			enemy.Health.Current -= s.Damage
			if enemy.Health.Current < 0 {
				enemy.Health.Current = 0
			}
			enemy.CurrentAttackCooldown = s.StunDuration
		}
	}
	
	return nil
}

// ChargeAbility rushes forward damaging enemies
type ChargeAbility struct {
	BaseAbility
	Damage      int
	Speed       float64
	Distance    float64
}

// NewChargeAbility creates a new charge ability
func NewChargeAbility() *ChargeAbility {
	return &ChargeAbility{
		BaseAbility: BaseAbility{
			Name:     "Charge",
			Cooldown: 180,
			ManaCost: 20,
		},
		Damage:   30,
		Speed:    400,
		Distance: 200,
	}
}

func (c *ChargeAbility) Execute(hero *entities.Hero, g GameInterface) error {
	if !c.IsReady() {
		return errors.New("ability on cooldown")
	}
	if hero.Mana.Current < c.ManaCost {
		return errors.New("not enough mana")
	}
	
	hero.Mana.Current -= c.ManaCost
	c.CurrentCooldown = c.Cooldown
	return nil
}

// ArmorUpPassive increases defense
type ArmorUpPassive struct {
	BaseAbility
	DefenseBonus int
}

// NewArmorUpPassive creates a new armor up passive
func NewArmorUpPassive() *ArmorUpPassive {
	return &ArmorUpPassive{
		BaseAbility: BaseAbility{
			Name: "Armor Up",
		},
		DefenseBonus: 5,
	}
}

func (a *ArmorUpPassive) Execute(hero *entities.Hero, g GameInterface) error {
	return nil
}

type BerserkerRageAbility struct {
	BaseAbility
	Duration      int
	AttackSpeedMultiplier float64
}

// NewBerserkerRageAbility creates a new berserker rage ability
func NewBerserkerRageAbility() *BerserkerRageAbility {
	return &BerserkerRageAbility{
		BaseAbility: BaseAbility{
			Name:     "Berserker Rage",
			Cooldown: 600,
			ManaCost: 0,
		},
		Duration:             480,
		AttackSpeedMultiplier: 3.0,
	}
}

func (b *BerserkerRageAbility) Execute(hero *entities.Hero, g GameInterface) error {
	return nil
}

