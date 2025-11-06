package systems

import (
	"game/internal/entities"
	"game/internal/events"
	"math/rand"
)

// CombatSystem handles all combat-related logic
type CombatSystem struct {
	listeners []events.EventListener
}

// NewCombatSystem creates a new combat system
func NewCombatSystem() *CombatSystem {
	return &CombatSystem{
		listeners: make([]events.EventListener, 0),
	}
}

// Subscribe adds an event listener
func (cs *CombatSystem) Subscribe(listener events.EventListener) {
	cs.listeners = append(cs.listeners, listener)
}

func (cs *CombatSystem) DamageEntity(target *entities.Entity, damage int, source *entities.Entity) {
	if target.Health == nil {
		return
	}
	
	oldHealth := target.Health.Current
	
	actualDamage := damage
	if target.Health.Current > 0 {
		variance := rand.Float64()*0.2 - 0.1
		actualDamage = int(float64(damage) * (1.0 + variance))
		
		if actualDamage < 1 {
			actualDamage = 1
		}
		
		target.Health.Current -= actualDamage
		if target.Health.Current < 0 {
			target.Health.Current = 0
		}
	}
	
	for _, listener := range cs.listeners {
		listener.OnHealthChanged(target, oldHealth, target.Health.Current)
	}
}

func (cs *CombatSystem) CalculateDamage(attacker *entities.Hero, target *entities.Entity) int {
	if target.Health == nil {
		return 0
	}
	
	baseDamage := attacker.Attack
	finalDamage := baseDamage - 2
	
	if finalDamage < 1 {
		finalDamage = 1
	}
	
	return finalDamage
}

// ApplyAbilityDamage applies damage from an ability
func (cs *CombatSystem) ApplyAbilityDamage(target *entities.Entity, damage int, source *entities.Entity) {
	cs.DamageEntity(target, damage, source)
}

