package entities

// Effect represents a temporary effect in the game
type Effect struct {
	Entity
	Type       string
	Duration   int
	Target     *Hero
}

