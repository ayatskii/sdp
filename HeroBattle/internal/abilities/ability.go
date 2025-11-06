package abilities

import "game/internal/entities"

// Ability is an alias for AbilityInterface to maintain compatibility
type Ability = entities.AbilityInterface

// GameInterface is an alias for entities.GameInterface
type GameInterface = entities.GameInterface

// BaseAbility provides common ability functionality
type BaseAbility struct {
	Name             string
	Cooldown         int
	CurrentCooldown  int
	ManaCost         int
}

// Update updates the cooldown
func (b *BaseAbility) Update(deltaTime float64) {
	if b.CurrentCooldown > 0 {
		b.CurrentCooldown--
	}
}

// IsReady checks if the ability is ready to use
func (b *BaseAbility) IsReady() bool {
	return b.CurrentCooldown <= 0
}

// GetName returns the ability name
func (b *BaseAbility) GetName() string {
	return b.Name
}

// GetCooldown returns the cooldown duration
func (b *BaseAbility) GetCooldown() int {
	return b.Cooldown
}

// GetCurrentCooldown returns the current cooldown
func (b *BaseAbility) GetCurrentCooldown() int {
	return b.CurrentCooldown
}

// GetManaCost returns the mana cost
func (b *BaseAbility) GetManaCost() int {
	return b.ManaCost
}

