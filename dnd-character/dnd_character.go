package dndcharacter

import (
	"math"
	"math/rand"
)

type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}

// Modifier calculates the ability modifier for a given ability score
func Modifier(score int) int {
	return int(math.Floor((float64(score) - 10) / 2))
}

// Ability uses randomness to generate the score for an ability
func Ability() (ability int) {
	min := 6

	for i := 0; i < 4; i++ {
		roll := rand.Intn(6) + 1
		ability += roll

		if roll < min {
			min = roll
		}
	}

	return ability - min
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
	c := Character{
		Ability(),
		Ability(),
		Ability(),
		Ability(),
		Ability(),
		Ability(),
		10,
	}

	c.Hitpoints += Modifier(c.Constitution)

	return c
}
