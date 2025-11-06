package input

import (
	"game/internal/abilities"
	"game/internal/entities"
)

// Command defines the interface for game commands
type Command interface {
	Execute(hero *entities.Hero, game abilities.GameInterface) error
}

type MoveCommand struct {
	Direction int
	Speed     float64
}

func (m *MoveCommand) Execute(hero *entities.Hero, game abilities.GameInterface) error {
	if hero == nil {
		return nil
	}
	
	hero.VelX, hero.VelY = 0, 0
	switch m.Direction {
	case 0:
		hero.VelY = -m.Speed
	case 1:
		hero.VelY = m.Speed
	case 2:
		hero.VelX = -m.Speed
	case 3:
		hero.VelX = m.Speed
	}
	return nil
}

type SwitchHeroCommand struct {
	Offset int
}

func (s *SwitchHeroCommand) Execute(hero *entities.Hero, game abilities.GameInterface) error {
	return nil
}

type CastAbilityCommand struct {
	AbilityIndex int
}

func (c *CastAbilityCommand) Execute(hero *entities.Hero, game abilities.GameInterface) error {
	if hero == nil {
		return nil
	}
	return hero.CastAbility(c.AbilityIndex, game)
}

type CastUltimateCommand struct{}

func (c *CastUltimateCommand) Execute(hero *entities.Hero, game abilities.GameInterface) error {
	if hero == nil {
		return nil
	}
	return hero.CastUltimate(game)
}

