package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/market-place/main_dadu/domain"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("SETUP GAME :")

	fmt.Println("Masukan Jumlah Pemain :")
	inputPlayerCMD, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Maaf permainan tidak dapat dimulai : %s", err)
	}
	inputPlayerCMD = strings.TrimSuffix(inputPlayerCMD, "\n")

	numberOfPlayer, err := strconv.Atoi(inputPlayerCMD)
	if err != nil {
		log.Fatalln("Maaf, jumlah player haruslah angka!")
	}

	fmt.Println("Masukan Jumlah Dice:")
	inputDiceCMD, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Maaf permainan tidak dapat dimulai : %s", err)
	}
	inputDiceCMD = strings.TrimSuffix(inputDiceCMD, "\n")

	numberOfDice, err := strconv.Atoi(inputDiceCMD)
	if err != nil {
		log.Fatalln("Maaf, jumlah player haruslah angka!")
	}

	game := domain.NewGame(numberOfPlayer, numberOfDice)
	for {
		if game.IsGameFinish() {
			fmt.Printf("The winner of game : player %d \n", game.GetWinner().ID)
			fmt.Printf("The loser of game : player %d \n", game.GetLoser().ID)
			break
		}

		//1.player shake dice
		game.PlayerShakeDice()

		time.Sleep(2 * time.Second)

		// //2. evaluate dice
		game.EvaluatePlayerDice()
		// for _, player := range game.Players {
		// 	fmt.Printf("Player : %d \n", player.ID)
		// 	for _, dice := range player.Dices {
		// 		fmt.Println(dice.Value)
		// 	}
		// }
	}
}
