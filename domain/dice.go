package domain

import (
	"math/rand"
	"time"
)

var (
	lastDiceID int = 0
)

type Dice struct {
	Value int
	ID    int
}

func NewDice() *Dice {

	// not conccurent safe
	dice := &Dice{
		Value: 0,
		ID:    lastDiceID + 1,
	}

	lastDiceID = lastDiceID + 1

	return dice
}

func (dice *Dice) Shake() int {
	rand.Seed(time.Now().UnixNano())
	max := 7
	min := 1
	random := rand.Intn(max - min)
	return random + min
}
