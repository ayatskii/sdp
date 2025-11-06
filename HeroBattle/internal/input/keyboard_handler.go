package input

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type KeyboardHandler struct {
	speed float64
}

func NewKeyboardHandler() *KeyboardHandler {
	return &KeyboardHandler{
		speed: 200.0 / 60.0,
	}
}

func (k *KeyboardHandler) GetMovementCommands() []Command {
	commands := make([]Command, 0)
	
	if IsKeyPressed(ebiten.KeyW) {
		commands = append(commands, &MoveCommand{Direction: 0, Speed: k.speed})
	}
	if IsKeyPressed(ebiten.KeyS) {
		commands = append(commands, &MoveCommand{Direction: 1, Speed: k.speed})
	}
	if IsKeyPressed(ebiten.KeyA) {
		commands = append(commands, &MoveCommand{Direction: 2, Speed: k.speed})
	}
	if IsKeyPressed(ebiten.KeyD) {
		commands = append(commands, &MoveCommand{Direction: 3, Speed: k.speed})
	}
	
	return commands
}

func (k *KeyboardHandler) GetJustPressedCommand() Command {
	if IsKeyJustPressed(ebiten.KeyQ) {
		return &SwitchHeroCommand{Offset: -1}
	}
	if IsKeyJustPressed(ebiten.KeyE) {
		return &SwitchHeroCommand{Offset: 1}
	}
	if IsKeyJustPressed(ebiten.Key1) {
		return &CastAbilityCommand{AbilityIndex: 0}
	}
	if IsKeyJustPressed(ebiten.Key2) {
		return &CastAbilityCommand{AbilityIndex: 1}
	}
	if IsKeyJustPressed(ebiten.Key3) {
		return &CastAbilityCommand{AbilityIndex: 2}
	}
	if IsKeyJustPressed(ebiten.Key5) {
		return &CastUltimateCommand{}
	}
	
	return nil
}

