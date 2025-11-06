package ui

import (
	"game/internal/entities"
	"game/internal/systems"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

// UISystem handles all UI rendering
type UISystem struct{}

// NewUISystem creates a new UI system
func NewUISystem() *UISystem {
	return &UISystem{}
}

func DrawHealthBar(screen *ebiten.Image, x, y, width, height float64, current, max int) {
	if max <= 0 {
		return
	}
	
	bgColor := color.RGBA{100, 0, 0, 255}
	drawRect(screen, x, y, width, height, bgColor)
	
	healthPercent := float64(current) / float64(max)
	if healthPercent > 1 {
		healthPercent = 1
	}
	if healthPercent < 0 {
		healthPercent = 0
	}
	
	healthWidth := width * healthPercent
	healthColor := color.RGBA{0, 200, 0, 255}
	if healthPercent < 0.3 {
		healthColor = color.RGBA{200, 0, 0, 255}
	} else if healthPercent < 0.6 {
		healthColor = color.RGBA{200, 200, 0, 255}
	}
	
	drawRect(screen, x, y, healthWidth, height, healthColor)
	
	borderColor := color.RGBA{255, 255, 255, 255}
	drawBorder(screen, x-1, y-1, width+2, height+2, borderColor)
}

func DrawManaBar(screen *ebiten.Image, x, y, width, height float64, current, max int) {
	if max <= 0 {
		return
	}
	
	bgColor := color.RGBA{0, 0, 100, 255}
	drawRect(screen, x, y, width, height, bgColor)
	
	manaPercent := float64(current) / float64(max)
	if manaPercent > 1 {
		manaPercent = 1
	}
	if manaPercent < 0 {
		manaPercent = 0
	}
	
	manaWidth := width * manaPercent
	manaColor := color.RGBA{100, 150, 255, 255}
	drawRect(screen, x, y, manaWidth, height, manaColor)
	
	borderColor := color.RGBA{255, 255, 255, 255}
	drawBorder(screen, x-1, y-1, width+2, height+2, borderColor)
}

func DrawGameUI(screen *ebiten.Image, currentHero *entities.Hero, heroes []*entities.Hero, enemies []*entities.Enemy, waveManager *systems.WaveManager, screenWidth int) {
	if currentHero == nil {
		return
	}
	
	panelX := 10.0
	panelY := 10.0
	
	if currentHero.Health != nil {
		DrawHealthBar(screen, panelX, panelY, 200, 20, currentHero.Health.Current, currentHero.Health.Max)
	}
	
	if currentHero.Mana.Max > 0 {
		DrawManaBar(screen, panelX, panelY+25, 200, 15, currentHero.Mana.Current, currentHero.Mana.Max)
	}
	
	ultimatePercent := float64(currentHero.UltimateEnergy) / 100.0
	if ultimatePercent > 1 {
		ultimatePercent = 1
	}
	if ultimatePercent < 0 {
		ultimatePercent = 0
	}
	ultimateWidth := 200.0 * ultimatePercent
	ultimateColor := color.RGBA{255, 200, 0, 255}
	drawRect(screen, panelX, panelY+45, 200, 10, color.RGBA{50, 50, 50, 255})
	if ultimateWidth > 0 {
		drawRect(screen, panelX, panelY+45, ultimateWidth, 10, ultimateColor)
	}
	
	abilityY := panelY + 60
	for i, ability := range currentHero.Abilities {
		if i >= 3 {
			break
		}
		cooldownX := panelX + float64(i*70)
		
		boxColor := color.RGBA{100, 100, 100, 255}
		if ability.IsReady() && currentHero.Mana.Current >= ability.GetManaCost() {
			boxColor = color.RGBA{150, 150, 150, 255}
		}
		drawRect(screen, cooldownX, abilityY, 60, 40, boxColor)
		
		if !ability.IsReady() {
			cooldownPercent := float64(ability.GetCurrentCooldown()) / float64(ability.GetCooldown())
			if cooldownPercent < 0 {
				cooldownPercent = 0
			}
			if cooldownPercent > 1 {
				cooldownPercent = 1
			}
			overlayHeight := 40.0 * cooldownPercent
			if overlayHeight > 0 {
				overlayColor := color.RGBA{0, 0, 0, 180}
				drawRect(screen, cooldownX, abilityY+40-overlayHeight, 60, overlayHeight, overlayColor)
			}
		}
		
		// Ability number
		// Draw simple text indicator (1, 2, 3)
		drawRect(screen, cooldownX+25, abilityY+15, 10, 10, color.RGBA{255, 255, 255, 255})
	}
	
	// Wave counter (top right)
	waveX := float64(screenWidth - 200)
	if waveX < 0 {
		waveX = 10 // Fallback if screen is too small
	}
	waveY := 10.0
	// Simple wave indicator (colored box)
	waveColor := color.RGBA{100, 100, 255, 255}
	drawRect(screen, waveX, waveY, 150, 30, waveColor)
	
	// Enemy count
	enemyCount := len(enemies)
	// Simple indicator box
	enemyColor := color.RGBA{255, 100, 100, 255}
	if enemyCount == 0 {
		enemyColor = color.RGBA{100, 255, 100, 255} // Green when no enemies
	}
	drawRect(screen, waveX, waveY+35, 150, 20, enemyColor)
}

func DrawEntityHealthBars(screen *ebiten.Image, enemies []*entities.Enemy, heroes []*entities.Hero) {
	for _, enemy := range enemies {
		if enemy.Health != nil && enemy.Health.Current > 0 {
			barX := enemy.X
			barY := enemy.Y - 10
			DrawHealthBar(screen, barX, barY, enemy.Width, 5, enemy.Health.Current, enemy.Health.Max)
		}
	}
	
	for _, hero := range heroes {
		if hero.Health != nil && hero.Health.Current > 0 {
			barX := hero.X
			barY := hero.Y - 10
			DrawHealthBar(screen, barX, barY, hero.Width, 5, hero.Health.Current, hero.Health.Max)
		}
	}
}

func drawRect(screen *ebiten.Image, x, y, w, h float64, c color.Color) {
	width := int(w)
	height := int(h)
	if width <= 0 || height <= 0 {
		return
	}
	
	img := ebiten.NewImage(width, height)
	img.Fill(c)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(x, y)
	screen.DrawImage(img, op)
}

func drawBorder(screen *ebiten.Image, x, y, w, h float64, c color.Color) {
	drawRect(screen, x, y, w, 1, c)
	drawRect(screen, x, y+h-1, w, 1, c)
	drawRect(screen, x, y, 1, h, c)
	drawRect(screen, x+w-1, y, 1, h, c)
}

