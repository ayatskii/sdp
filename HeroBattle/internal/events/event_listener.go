package events

import "game/internal/entities"

// EventListener defines the interface for event listeners
type EventListener interface {
	OnHealthChanged(entity *entities.Entity, oldHealth, newHealth int)
	OnCooldownUpdated(ability entities.AbilityInterface)
	OnEnemyDied(enemy *entities.Enemy)
	OnWaveCompleted()
	OnAbilityUsed(hero *entities.Hero, abilityName string)
}

