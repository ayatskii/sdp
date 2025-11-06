package entities

// AbilityInterface defines the interface for abilities
// This avoids circular dependencies
type AbilityInterface interface {
	Execute(hero *Hero, game GameInterface) error
	Update(deltaTime float64)
	IsReady() bool
	GetName() string
	GetCooldown() int
	GetCurrentCooldown() int
	GetManaCost() int
}

// GameInterface defines what abilities can do with the game
type GameInterface interface {
	AddProjectile(projectile *Projectile)
	GetEnemies() []*Enemy
	AddEffect(effect *Effect)
}

