package domain

import "fmt"

const (
	DICE_NUMBER_TO_REMOVE     int = 6
	DICE_NUMBER_TO_THROW_NEXT int = 1
)

type gamePlay struct {
	Players []*player
}

func NewGame(numberOfPlayer int, numberOfDice int) *gamePlay {
	players := []*player{}

	for i := 0; i < numberOfPlayer; i++ {
		player := NewPlayer(numberOfDice)
		players = append(players, player)
	}

	return &gamePlay{
		Players: players,
	}
}

func (game *gamePlay) IsGameFinish() bool {
	numberOfFinishPlayer := 0
	for _, player := range game.Players {
		if player.IsPlayerWin() {
			numberOfFinishPlayer++
		}
	}

	return numberOfFinishPlayer == (len(game.Players) - 1)
}

func (game *gamePlay) PlayerShakeDice() {
	for _, player := range game.Players {
		if !player.IsPlayerWin() {
			player.ShakeDice()
		}
	}
}

func (game *gamePlay) EvaluatePlayerDice() {
	fmt.Println("EVALUATE DICE!")
	type History struct {
		From *player
		To   *player
		Dice *Dice
	}
	historyThrow := []History{}

	for indexPlayer, player := range game.Players {
		for _, dice := range player.Dices {
			//remove dice from player if dice's value is 6
			if dice.Value == DICE_NUMBER_TO_REMOVE {
				player.RemoveDice(dice.ID)
			}

			//throw dice to next player if dice's value is 1
			if dice.Value == DICE_NUMBER_TO_THROW_NEXT {
				nextPlayerIndex := game.getNextPlayerWhoStillPlay(indexPlayer + 1)
				if nextPlayerIndex != -1 {
					nextPlayer := game.Players[nextPlayerIndex]
					historyThrow = append(historyThrow, History{
						From: player,
						To:   nextPlayer,
						Dice: dice,
					})
				}
			}
		}
	}

	for _, history := range historyThrow {
		history.From.ThrowDiceToNextPlayer(history.Dice, history.To)
	}

	game.showPlayerDice()
}

func (game *gamePlay) GetWinner() *player {
	maxScore := 0
	playerIndex := -1

	for indexPlayer, player := range game.Players {
		if player.Point > maxScore {
			maxScore = player.Point
			playerIndex = indexPlayer
		}
	}

	return game.Players[playerIndex]
}

func (game *gamePlay) GetLoser() *player {
	for _, player := range game.Players {
		if !player.IsPlayerWin() {
			return player
		}
	}

	return nil
}

func (game *gamePlay) showPlayerDice() {
	for _, player := range game.Players {

		if player.IsPlayerWin() {
			fmt.Printf("Player %d has end the game!\n", player.ID)
		} else {
			fmt.Printf("Player  %d dices in hand: \n", player.ID)

			for _, dice := range player.Dices {
				fmt.Printf("\t dice id : %d - dice value : %d \n", dice.ID, dice.Value)
			}
		}
	}
}

func (game *gamePlay) getNextPlayerWhoStillPlay(startIndex int) int {
	for i := startIndex; i < len(game.Players); i++ {
		player := game.Players[i]
		if !player.IsPlayerWin() {
			return i
		}
	}

	for i := 0; i < startIndex; i++ {
		player := game.Players[i]
		if !player.IsPlayerWin() {
			return i
		}
	}

	return -1
}
