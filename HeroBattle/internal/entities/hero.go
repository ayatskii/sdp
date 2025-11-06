package entities

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

// Hero represents a playable hero character
type Hero struct {
	Entity
	Name          string
	Mana          struct {
		Current int
		Max     int
	}
	Attack        int
	Defense       int
	Abilities     []AbilityInterface
	Passive       AbilityInterface
	Ultimate      AbilityInterface
	UltimateEnergy int
}

// Update updates the hero
func (h *Hero) Update() {
	h.Entity.Update()
	
	for _, ability := range h.Abilities {
		ability.Update(1.0 / 60.0)
	}
	if h.Passive != nil {
		h.Passive.Update(1.0 / 60.0)
		h.applyPassiveEffects()
	}
	if h.Ultimate != nil {
		h.Ultimate.Update(1.0 / 60.0)
	}
}

func (h *Hero) applyPassiveEffects() {
	if h.Passive.GetName() == "Mana Regeneration" {
		h.Mana.Current += 1
		if h.Mana.Current > h.Mana.Max {
			h.Mana.Current = h.Mana.Max
		}
	}
}

// Draw draws the hero
func (h *Hero) Draw(screen *ebiten.Image) {
	h.Entity.Draw(screen)
}

// CastAbility casts an ability at the given index
func (h *Hero) CastAbility(abilityIndex int, g GameInterface) error {
	if abilityIndex < 0 || abilityIndex >= len(h.Abilities) {
		return nil
	}
	
	ability := h.Abilities[abilityIndex]
	if err := ability.Execute(h, g); err != nil {
		return err
	}
	
	h.UltimateEnergy += 5
	if h.UltimateEnergy > 100 {
		h.UltimateEnergy = 100
	}
	
	return nil
}

// CastUltimate casts the ultimate ability
func (h *Hero) CastUltimate(g GameInterface) error {
	if h.UltimateEnergy < 100 {
		return nil
	}
	
	if h.Ultimate != nil {
		if err := h.Ultimate.Execute(h, g); err != nil {
			return err
		}
		h.UltimateEnergy = 0
	}
	return nil
}

// NewHero creates a new hero with default values
func NewHero(name string, x, y float64) *Hero {
	return &Hero{
		Entity: Entity{
			ID:     name,
			X:      x,
			Y:      y,
			Width:  32,
			Height: 32,
			Health: &Health{
				Current: 100,
				Max:     100,
			},
		},
		Name: name,
		Mana: struct {
			Current int
			Max     int
		}{
			Current: 50,
			Max:     50,
		},
		Attack:        10,
		Defense:       5,
		Abilities:     make([]AbilityInterface, 0),
		UltimateEnergy: 0,
	}
}

// SetSprite sets the hero's sprite
func (h *Hero) SetSprite(c color.Color) {
	h.Sprite = NewSpriteFromColor(int(h.Width), int(h.Height), c)
}

