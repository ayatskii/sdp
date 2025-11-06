package systems

import (
	"game/internal/entities"
	"math"
)

// EnemyAI handles enemy AI behavior
type EnemyAI struct{}

// NewEnemyAI creates a new enemy AI system
func NewEnemyAI() *EnemyAI {
	return &EnemyAI{}
}

func (ai *EnemyAI) UpdateEnemy(enemy *entities.Enemy, hero *entities.Hero) {
	if hero == nil || hero.Health == nil || hero.Health.Current <= 0 {
		ai.patrolRandomly(enemy)
		return
	}
	
	distToHero := entities.GetDistance(enemy.X, enemy.Y, hero.X, hero.Y)
	
	if distToHero > enemy.ViewRange {
		ai.patrolRandomly(enemy)
	} else if distToHero > enemy.AttackRange {
		ai.moveToward(enemy, hero.X, hero.Y)
	} else {
		enemy.VelX = 0
		enemy.VelY = 0
	}
}

func (ai *EnemyAI) moveToward(enemy *entities.Enemy, targetX, targetY float64) {
	dx := targetX - enemy.X
	dy := targetY - enemy.Y
	distance := math.Sqrt(dx*dx + dy*dy)
	
	if distance > 0 {
		// Normalize direction
		enemy.VelX = (dx / distance) * enemy.Speed
		enemy.VelY = (dy / distance) * enemy.Speed
	}
}

func (ai *EnemyAI) patrolRandomly(enemy *entities.Enemy) {
	enemy.VelX = 0
	enemy.VelY = 0
}

