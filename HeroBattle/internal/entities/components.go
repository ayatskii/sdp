package entities

import (
	"game/internal/assets"
	
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

// Sprite represents a visual component
type Sprite struct {
	Image  *ebiten.Image
	FrameX int
	FrameY int
	Color  color.Color
}

// Health represents health statistics
type Health struct {
	Current int
	Max     int
}

// NewSpriteFromColor creates a colored rectangle sprite (fallback)
func NewSpriteFromColor(width, height int, c color.Color) *Sprite {
	img := ebiten.NewImage(width, height)
	img.Fill(c)
	return &Sprite{
		Image:  img,
		Color:  c,
		FrameX: width,
		FrameY: height,
	}
}

// NewSpriteFromAsset loads a sprite from assets
func NewSpriteFromAsset(assetKey string) *Sprite {
	img := assets.GetImage(assetKey)
	if img == nil {
		// Fallback to colored rectangle if asset not found
		return NewSpriteFromColor(32, 32, color.RGBA{128, 128, 128, 255})
	}
	
	bounds := img.Bounds()
	return &Sprite{
		Image:  img,
		FrameX: bounds.Dx(),
		FrameY: bounds.Dy(),
	}
}

