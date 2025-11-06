package input

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// InputHandler handles keyboard input
type InputHandler struct{}

// NewInputHandler creates a new input handler
func NewInputHandler() *InputHandler {
	return &InputHandler{}
}

// IsKeyPressed checks if a key is currently pressed
func IsKeyPressed(key ebiten.Key) bool {
	return ebiten.IsKeyPressed(key)
}

// IsKeyJustPressed checks if a key was just pressed this frame
func IsKeyJustPressed(key ebiten.Key) bool {
	return inpututil.IsKeyJustPressed(key)
}

// GetCursorPosition returns the cursor position
func GetCursorPosition() (int, int) {
	return ebiten.CursorPosition()
}

