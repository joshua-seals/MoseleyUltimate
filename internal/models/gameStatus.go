package models

import (
	"sync"
	"time"
)

type GameStatus struct {
	defaultStatus string
	override      string
	overrideUntil time.Time
	mu            sync.Mutex
}

func NewGameStatus() *GameStatus {
	return &GameStatus{
		defaultStatus: "GAME ON",
	}
}

// GetStatus returns the current status, reverting after expiration
func GetStatus() string {
	gameStatus.mu.Lock()
	defer gameStatus.mu.Unlock()

	if gameStatus.override != "" && time.Now().Before(gameStatus.overrideUntil) {
		return gameStatus.override
	}
	return gameStatus.defaultStatus
}

// SetTemporaryOverride sets an override that expires after a duration
func SetTemporaryOverride(status string, duration time.Duration) {
	gameStatus.mu.Lock()
	defer gameStatus.mu.Unlock()

	gameStatus.override = status
	gameStatus.overrideUntil = time.Now().Add(duration)
}

// Instantiate the variable
var gameStatus = NewGameStatus()
