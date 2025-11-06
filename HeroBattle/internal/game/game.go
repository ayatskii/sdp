package game

import (
	"game/internal/entities"
	"game/internal/factories"
	"game/internal/input"
	"game/internal/systems"
	"game/internal/ui"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

// Game represents the main game state
type Game struct {
	entities         []entities.EntityInterface
	heroes           []*entities.Hero
	currentHeroIndex int
	enemies          []*entities.Enemy
	projectiles      []*entities.Projectile
	effects          []*entities.Effect
	combatSystem     *systems.CombatSystem
	enemyAI          *systems.EnemyAI
	waveManager      *systems.WaveManager
	gameState        string
}

var _ entities.GameInterface = (*Game)(nil)

// NewGame creates a new game instance
func NewGame() *Game {
	factory := &factories.HeroFactory{}

	heroes := []*entities.Hero{
		factory.CreateMage(100, 300),
		factory.CreateKnight(200, 300),
		factory.CreateArcher(300, 300),
	}

	g := &Game{
		entities:         make([]entities.EntityInterface, 0),
		heroes:           heroes,
		enemies:          make([]*entities.Enemy, 0),
		projectiles:      make([]*entities.Projectile, 0),
		effects:          make([]*entities.Effect, 0),
		currentHeroIndex: 0,
		combatSystem:     systems.NewCombatSystem(),
		enemyAI:          systems.NewEnemyAI(),
		waveManager:      systems.NewWaveManager(),
		gameState:        "playing",
	}

	for _, hero := range heroes {
		g.entities = append(g.entities, hero)
	}

	return g
}

// AddProjectile adds a projectile to the game
func (g *Game) AddProjectile(projectile *entities.Projectile) {
	g.projectiles = append(g.projectiles, projectile)
	g.entities = append(g.entities, projectile)
}

// GetEnemies returns all enemies
func (g *Game) GetEnemies() []*entities.Enemy {
	return g.enemies
}

// AddEnemy adds an enemy to the game
func (g *Game) AddEnemy(enemy *entities.Enemy) {
	g.enemies = append(g.enemies, enemy)
	g.entities = append(g.entities, enemy)
}

// GetEnemyCount returns the number of enemies
func (g *Game) GetEnemyCount() int {
	return len(g.enemies)
}

// AddEffect adds an effect to the game
func (g *Game) AddEffect(effect *entities.Effect) {
	g.effects = append(g.effects, effect)
	g.entities = append(g.entities, effect)
}

// GetCurrentHero returns the currently active hero
func (g *Game) GetCurrentHero() *entities.Hero {
	if len(g.heroes) == 0 {
		return nil
	}
	return g.heroes[g.currentHeroIndex]
}

// SwitchToHero switches to the next or previous hero
func (g *Game) SwitchToHero(offset int) {
	if len(g.heroes) == 0 {
		return
	}
	g.currentHeroIndex = (g.currentHeroIndex + offset + len(g.heroes)) % len(g.heroes)
}

// Update updates the game state
func (g *Game) Update() error {
	if g.gameState != "playing" {
		// Game over, just check for restart
		if input.IsKeyJustPressed(ebiten.KeyR) {
			// Restart game (would need to reset state)
		}
		return nil
	}

	// Handle input using Bridge pattern (Command pattern)
	keyboardHandler := input.NewKeyboardHandler()
	currentHero := g.GetCurrentHero()

	if cmd := keyboardHandler.GetJustPressedCommand(); cmd != nil {
		switch c := cmd.(type) {
		case *input.SwitchHeroCommand:
			g.SwitchToHero(c.Offset)
		case *input.CastAbilityCommand:
			if currentHero != nil && c.AbilityIndex >= 0 && c.AbilityIndex < len(currentHero.Abilities) {
				currentHero.CastAbility(c.AbilityIndex, g)
			}
		case *input.CastUltimateCommand:
			if currentHero != nil {
				currentHero.CastUltimate(g)
			}
		}
	}

	if currentHero != nil {
		currentHero.VelX = 0
		currentHero.VelY = 0

		movementCommands := keyboardHandler.GetMovementCommands()
		for _, moveCmd := range movementCommands {
			moveCmd.Execute(currentHero, g)
		}

		if currentHero.X < 0 {
			currentHero.X = 0
		}
		if currentHero.X+currentHero.Width > ScreenWidth {
			currentHero.X = ScreenWidth - currentHero.Width
		}
		if currentHero.Y < 0 {
			currentHero.Y = 0
		}
		if currentHero.Y+currentHero.Height > ScreenHeight {
			currentHero.Y = ScreenHeight - currentHero.Height
		}
	}

	if g.waveManager.IsWaveComplete(g.GetEnemyCount) && g.waveManager.GetCurrentWave() < g.waveManager.GetMaxWaves() {
	}
	g.waveManager.Update(g.AddEnemy, ScreenWidth, ScreenHeight)

	for _, enemy := range g.enemies {
		if currentHero != nil {
			g.enemyAI.UpdateEnemy(enemy, currentHero)
		}
	}

	for _, entity := range g.entities {
		entity.Update()
	}

	g.updateEnemyAttacks()
	g.updateCollisions()
	g.checkGameState()

	return nil
}

func (g *Game) updateEnemyAttacks() {
	currentHero := g.GetCurrentHero()
	if currentHero == nil || currentHero.Health == nil {
		return
	}

	for _, enemy := range g.enemies {
		if enemy.Health == nil || enemy.Health.Current <= 0 {
			continue
		}

		dist := entities.GetDistance(enemy.X, enemy.Y, currentHero.X, currentHero.Y)
		if dist <= enemy.AttackRange && enemy.CurrentAttackCooldown <= 0 {
			damage := enemy.Attack
			g.combatSystem.DamageEntity(&currentHero.Entity, damage, &enemy.Entity)
			enemy.CurrentAttackCooldown = enemy.AttackCooldown
		}
	}
}

func (g *Game) updateCollisions() {
	aliveProjectiles := make([]*entities.Projectile, 0)
	for _, proj := range g.projectiles {
		if proj.IsExpired() {
			for i, entity := range g.entities {
				if entity == proj {
					g.entities = append(g.entities[:i], g.entities[i+1:]...)
					break
				}
			}
			continue
		}

		projCenterX := proj.X + proj.Width/2
		projCenterY := proj.Y + proj.Height/2

		hit := false
		for _, enemy := range g.enemies {
			if enemy.Health == nil || enemy.Health.Current <= 0 {
				continue
			}

			enemyCenterX := enemy.X + enemy.Width/2
			enemyCenterY := enemy.Y + enemy.Height/2
			dist := entities.GetDistance(projCenterX, projCenterY, enemyCenterX, enemyCenterY)

			if !proj.IsExplosive && dist < (enemy.Width/2+proj.Width/2) {
				enemy.Health.Current -= proj.Damage
				if enemy.Health.Current < 0 {
					enemy.Health.Current = 0
				}
				hit = true
				break
			} else if proj.IsExplosive && dist <= proj.AOERadius {
				enemy.Health.Current -= proj.Damage
				if enemy.Health.Current < 0 {
					enemy.Health.Current = 0
				}
				hit = true
			}
		}

		if hit || proj.IsExpired() {
			for i, entity := range g.entities {
				if entity == proj {
					g.entities = append(g.entities[:i], g.entities[i+1:]...)
					break
				}
			}
		} else {
			aliveProjectiles = append(aliveProjectiles, proj)
		}
	}
	g.projectiles = aliveProjectiles
	g.enemies = g.removeDeadEnemies(g.enemies)
}

func (g *Game) removeDeadEnemies(enemies []*entities.Enemy) []*entities.Enemy {
	alive := make([]*entities.Enemy, 0)
	for _, enemy := range enemies {
		if enemy.Health != nil && enemy.Health.Current > 0 {
			alive = append(alive, enemy)
		} else {
			for i, entity := range g.entities {
				if entity == enemy {
					g.entities = append(g.entities[:i], g.entities[i+1:]...)
					break
				}
			}
		}
	}
	return alive
}

func (g *Game) checkGameState() {
	allDead := true
	for _, hero := range g.heroes {
		if hero.Health != nil && hero.Health.Current > 0 {
			allDead = false
			break
		}
	}
	if allDead {
		g.gameState = "defeat"
		return
	}

	if g.waveManager.IsGameComplete() && len(g.enemies) == 0 {
		g.gameState = "victory"
		return
	}
}

// Draw draws the game
func (g *Game) Draw(screen *ebiten.Image) {
	// Clear screen with black background
	screen.Fill(color.RGBA{0, 0, 0, 255})

	for _, entity := range g.entities {
		entity.Draw(screen)
	}

	ui.DrawEntityHealthBars(screen, g.enemies, g.heroes)

	currentHero := g.GetCurrentHero()
	if currentHero != nil {
		drawBorder(screen, currentHero.X-2, currentHero.Y-2,
			currentHero.Width+4, currentHero.Height+4, color.RGBA{255, 255, 0, 255})
	}

	ui.DrawGameUI(screen, currentHero, g.heroes, g.enemies, g.waveManager, ScreenWidth)

	if g.gameState == "victory" {
		drawGameOverScreen(screen, "VICTORY!", color.RGBA{0, 255, 0, 255})
	} else if g.gameState == "defeat" {
		drawGameOverScreen(screen, "DEFEAT!", color.RGBA{255, 0, 0, 255})
	}
}

func drawGameOverScreen(screen *ebiten.Image, message string, c color.Color) {
	overlay := ebiten.NewImage(ScreenWidth, ScreenHeight)
	overlay.Fill(color.RGBA{0, 0, 0, 180})
	screen.DrawImage(overlay, nil)

	boxX := float64(ScreenWidth/2 - 150)
	boxY := float64(ScreenHeight/2 - 50)
	boxW := 300.0
	boxH := 100.0

	drawRect(screen, boxX, boxY, boxW, boxH, color.RGBA{50, 50, 50, 255})
	drawBorder(screen, boxX, boxY, boxW, boxH, c)
	drawRect(screen, boxX+10, boxY+40, boxW-20, 20, c)
}

func drawBorder(screen *ebiten.Image, x, y, w, h float64, c color.Color) {
	drawRect(screen, x, y, w, 2, c)
	drawRect(screen, x, y+h-2, w, 2, c)
	drawRect(screen, x, y, 2, h, c)
	drawRect(screen, x+w-2, y, 2, h, c)
}

func drawRect(screen *ebiten.Image, x, y, w, h float64, c color.Color) {
	img := ebiten.NewImage(int(w), int(h))
	img.Fill(c)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(x, y)
	screen.DrawImage(img, op)
}

// Layout returns the game's logical screen size
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}

// AddEntity adds an entity to the game
func (g *Game) AddEntity(entity entities.EntityInterface) {
	g.entities = append(g.entities, entity)
}

// GetEntities returns all entities
func (g *Game) GetEntities() []entities.EntityInterface {
	return g.entities
}

const (
	ScreenWidth  = 800
	ScreenHeight = 600
)
