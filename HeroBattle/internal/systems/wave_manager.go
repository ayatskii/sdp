package systems

import "game/internal/entities"

// WaveManager manages enemy waves
type WaveManager struct {
	currentWave int
	maxWaves    int
	enemiesSpawned int
	enemiesPerWave []int
	enemyTypes      [][]string
	spawnCooldown   int
	currentSpawnCooldown int
}

// NewWaveManager creates a new wave manager
func NewWaveManager() *WaveManager {
	return &WaveManager{
		currentWave: 0,
		maxWaves:    5,
		enemiesSpawned: 0,
		enemiesPerWave: []int{5, 5, 6, 7, 8},
		enemyTypes: [][]string{
			{"Melee", "Melee", "Melee", "Melee", "Melee"},
			{"Melee", "Melee", "Melee", "Ranged", "Ranged"},
			{"Melee", "Melee", "Ranged", "Ranged", "Tank"},
			{"Ranged", "Ranged", "Ranged", "Tank", "Tank"},
			{"Melee", "Melee", "Ranged", "Ranged", "Tank", "Tank", "Tank", "Tank"},
		},
		spawnCooldown: 60,
		currentSpawnCooldown: 0,
	}
}

func (wm *WaveManager) Update(addEnemy func(enemy *entities.Enemy), screenWidth, screenHeight int) {
	if wm.currentWave < wm.maxWaves && wm.enemiesSpawned == 0 {
		wm.currentWave++
		wm.enemiesSpawned = 0
	}
	
	if wm.currentWave > 0 && wm.currentWave <= wm.maxWaves {
		waveIndex := wm.currentWave - 1
		totalEnemies := wm.enemiesPerWave[waveIndex]
		
		if wm.enemiesSpawned < totalEnemies {
			if wm.currentSpawnCooldown <= 0 {
				waveIndex := wm.currentWave - 1
				enemyIndex := wm.enemiesSpawned
				if enemyIndex < len(wm.enemyTypes[waveIndex]) {
					enemyType := wm.enemyTypes[waveIndex][enemyIndex]
					wm.spawnEnemy(addEnemy, enemyType, screenWidth, screenHeight)
					wm.enemiesSpawned++
					wm.currentSpawnCooldown = wm.spawnCooldown
				}
			} else {
				wm.currentSpawnCooldown--
			}
		}
	}
}

func (wm *WaveManager) spawnEnemy(addEnemy func(enemy *entities.Enemy), enemyType string, screenWidth, screenHeight int) {
	x, y := float64(0), float64(0)
	
	edge := wm.enemiesSpawned % 4
	switch edge {
	case 0:
		x = float64(wm.enemiesSpawned * 100 % screenWidth)
		y = 0
	case 1:
		x = float64(screenWidth - 32)
		y = float64(wm.enemiesSpawned * 80 % screenHeight)
	case 2:
		x = float64(wm.enemiesSpawned * 100 % screenWidth)
		y = float64(screenHeight - 32)
	case 3:
		x = 0
		y = float64(wm.enemiesSpawned * 80 % screenHeight)
	}
	
	enemy := entities.NewEnemy(enemyType, x, y)
	addEnemy(enemy)
}

func (wm *WaveManager) IsWaveComplete(getEnemyCount func() int) bool {
	if wm.currentWave == 0 {
		return false
	}
	
	waveIndex := wm.currentWave - 1
	totalEnemies := wm.enemiesPerWave[waveIndex]
	
	if wm.enemiesSpawned < totalEnemies {
		return false
	}
	
	return getEnemyCount() == 0
}

// GetCurrentWave returns the current wave number
func (wm *WaveManager) GetCurrentWave() int {
	return wm.currentWave
}

// GetMaxWaves returns the maximum number of waves
func (wm *WaveManager) GetMaxWaves() int {
	return wm.maxWaves
}

// IsGameComplete checks if all waves are complete
func (wm *WaveManager) IsGameComplete() bool {
	return wm.currentWave > wm.maxWaves && wm.enemiesSpawned >= wm.enemiesPerWave[wm.maxWaves-1]
}

