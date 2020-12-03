package domain

import "fmt"

var (
	lastPlayerID int = 0
)

type player struct {
	ID    int
	Point int
	Dices []*Dice
}

func NewPlayer(numOfDice int) *player {

	dices := []*Dice{}
	for i := 0; i < numOfDice; i++ {
		dice := NewDice()
		dices = append(dices, dice)
	}

	p := &player{
		ID:    lastPlayerID + 1,
		Dices: dices,
	}

	lastPlayerID = lastPlayerID + 1

	return p
}

func (p *player) ShakeDice() {
	fmt.Printf("Player %d shakes dices! \n", p.ID)
	for _, dice := range p.Dices {
		dice.Value = dice.Shake()
		fmt.Printf("Dice with id %d : value : %d \n", dice.ID, dice.Value)
	}
}

func (p *player) ThrowDiceToNextPlayer(dice *Dice, nextPlayer *player) {
	//give to next player
	nextPlayer.Dices = append(nextPlayer.Dices, dice)

	//remove from our self
	p.spliceDice(dice.ID)
}

func (p *player) RemoveDice(diceID int) {
	fmt.Printf("Player %d : Remove Dice with id : %d \n", p.ID, diceID)
	p.spliceDice(diceID)
	p.Point++
}

func (p *player) IsPlayerWin() bool {
	return len(p.Dices) == 0
}

func (p *player) spliceDice(diceID int) {
	index := p.findIndexDice(diceID)
	if index != -1 {
		newDices := []*Dice{}
		isDiceLastItem := len(p.Dices) == 1
		if isDiceLastItem {
			p.Dices = newDices
			return
		}

		isDiceInFirst := index == 0
		if isDiceInFirst {
			p.Dices = append(newDices, p.Dices[index+1:]...)
			return
		}

		isDiceInLast := index == (len(p.Dices) - 1)
		if isDiceInLast {
			p.Dices = append(newDices, p.Dices[:index]...)
			return
		}

		isDiceInMidle := (index > 0) && (index < (len(p.Dices) - 1))
		if isDiceInMidle {
			p.Dices = append(p.Dices[:index], p.Dices[index+1:]...)
			return
		}
	}
}

func (p *player) findIndexDice(diceID int) int {
	for index, dice := range p.Dices {
		if dice.ID == diceID {
			return index
		}
	}

	return -1
}
