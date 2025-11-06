package factories

import (
	"game/internal/abilities"
	"game/internal/entities"

	"image/color"
)

// HeroFactory creates different hero types
type HeroFactory struct{}

// NewHeroFactory creates a new hero factory
func NewHeroFactory() *HeroFactory {
	return &HeroFactory{}
}

// CreateMage creates a Mage hero
func (f *HeroFactory) CreateMage(x, y float64) *entities.Hero {
	hero := entities.NewHero("Mage", x, y)
	hero.Health.Max = 60
	hero.Health.Current = 60
	hero.Mana.Max = 100
	hero.Mana.Current = 100
	hero.Attack = 8
	hero.Defense = 3
	
	hero.Abilities = []abilities.Ability{
		abilities.NewFireballAbility(),
		abilities.NewFrostNovaAbility(),
		abilities.NewArcaneShieldAbility(),
	}
	hero.Passive = abilities.NewManaRegenPassive()
	hero.Ultimate = abilities.NewMeteorStormAbility()
	
	hero.Sprite = entities.NewSpriteFromAsset("mage")
	if hero.Sprite == nil || hero.Sprite.Image == nil {
		hero.SetSprite(color.RGBA{100, 50, 200, 255})
	}
	return hero
}

// CreateKnight creates a Knight hero
func (f *HeroFactory) CreateKnight(x, y float64) *entities.Hero {
	hero := entities.NewHero("Knight", x, y)
	hero.Health.Max = 120
	hero.Health.Current = 120
	hero.Mana.Max = 40
	hero.Mana.Current = 40
	hero.Attack = 12
	hero.Defense = 8
	
	hero.Abilities = []abilities.Ability{
		abilities.NewSlashAbility(),
		abilities.NewShieldBashAbility(),
		abilities.NewChargeAbility(),
	}
	hero.Passive = abilities.NewArmorUpPassive()
	hero.Ultimate = abilities.NewBerserkerRageAbility()
	
	hero.Sprite = entities.NewSpriteFromAsset("knight")
	if hero.Sprite == nil || hero.Sprite.Image == nil {
		hero.SetSprite(color.RGBA{150, 150, 150, 255})
	}
	return hero
}

// CreateArcher creates an Archer hero
func (f *HeroFactory) CreateArcher(x, y float64) *entities.Hero {
	hero := entities.NewHero("Archer", x, y)
	hero.Health.Max = 80
	hero.Health.Current = 80
	hero.Mana.Max = 60
	hero.Mana.Current = 60
	hero.Attack = 10
	hero.Defense = 4
	
	hero.Abilities = []abilities.Ability{
		abilities.NewQuickShotAbility(),
		abilities.NewArrowRainAbility(),
		abilities.NewEvasionAbility(),
	}
	hero.Passive = abilities.NewCriticalStrikePassive()
	hero.Ultimate = abilities.NewPerfectShotAbility()
	
	hero.Sprite = entities.NewSpriteFromAsset("archer")
	if hero.Sprite == nil || hero.Sprite.Image == nil {
		hero.SetSprite(color.RGBA{150, 100, 50, 255})
	}
	return hero
}

