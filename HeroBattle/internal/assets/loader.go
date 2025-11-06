package assets

import (
	"image"
	_ "image/png"
	_ "image/jpeg"
	"os"
	"path/filepath"
	
	"github.com/hajimehoshi/ebiten/v2"
)

// AssetManager manages all game assets
type AssetManager struct {
	images     map[string]*ebiten.Image
	assetsPath string
	scaleFactor float64
}

var globalAssetManager *AssetManager

func InitAssets() error {
	return InitAssetsWithScale(1.0)
}

func InitAssetsWithScale(scaleFactor float64) error {
	globalAssetManager = &AssetManager{
		images:      make(map[string]*ebiten.Image),
		assetsPath:  "assets",
		scaleFactor: scaleFactor,
	}
	
	if _, err := os.Stat("assets"); os.IsNotExist(err) {
		if _, err := os.Stat("internal/assets"); err == nil {
			globalAssetManager.assetsPath = "internal/assets"
		}
	}
	
	return globalAssetManager.loadAll()
}

func (am *AssetManager) loadAll() error {
	am.LoadImage("sprites/heroes/mage.png", "mage")
	am.LoadImage("sprites/heroes/knight.png", "knight")
	am.LoadImage("sprites/heroes/archer.png", "archer")
	am.LoadImage("sprites/enemies/melee.png", "melee")
	am.LoadImage("sprites/enemies/ranged.png", "ranged")
	am.LoadImage("sprites/enemies/tank.png", "tank")
	am.LoadImage("sprites/projectiles/fireball.png", "fireball")
	am.LoadImage("sprites/projectiles/arrow.png", "arrow")
	return nil
}

func (am *AssetManager) LoadImage(path, key string) error {
	fullPath := filepath.Join(am.assetsPath, path)
	file, err := os.Open(fullPath)
	if err != nil {
		return nil
	}
	defer file.Close()
	
	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}
	
	bounds := img.Bounds()
	originalWidth := bounds.Dx()
	originalHeight := bounds.Dy()
	
	scaledWidth := int(float64(originalWidth) * am.scaleFactor)
	scaledHeight := int(float64(originalHeight) * am.scaleFactor)
	
	if scaledWidth < 1 {
		scaledWidth = 1
	}
	if scaledHeight < 1 {
		scaledHeight = 1
	}
	
	originalEbitenImage := ebiten.NewImageFromImage(img)
	
	if am.scaleFactor == 1.0 {
		am.images[key] = originalEbitenImage
		return nil
	}
	
	scaledImage := ebiten.NewImage(scaledWidth, scaledHeight)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(am.scaleFactor, am.scaleFactor)
	scaledImage.DrawImage(originalEbitenImage, op)
	
	am.images[key] = scaledImage
	return nil
}

// GetImage returns an image by key
func GetImage(key string) *ebiten.Image {
	if globalAssetManager == nil {
		return nil
	}
	return globalAssetManager.images[key]
}

// GetImageCopy returns a copy of an image
func GetImageCopy(key string) *ebiten.Image {
	img := GetImage(key)
	if img == nil {
		return nil
	}
	
	bounds := img.Bounds()
	newImg := ebiten.NewImage(bounds.Dx(), bounds.Dy())
	op := &ebiten.DrawImageOptions{}
	newImg.DrawImage(img, op)
	return newImg
}

// HasImage checks if an image exists
func HasImage(key string) bool {
	if globalAssetManager == nil {
		return false
	}
	_, exists := globalAssetManager.images[key]
	return exists
}

