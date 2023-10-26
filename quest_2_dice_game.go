package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

type Player struct {
	Number int
	Score  int
	Dice   []int
}

var players []Player

// Set player number
func (p *Player) SetNumber(number int) {
	p.Number = number
}

// Roll dice
func (p *Player) RollDice() int {
	// Random number between 1 and 6
	return rand.Intn(6) + 1
}

// Roll dices
func (p *Player) RollDices(dices int) []int {
	var d []int

	for i := 0; i < dices; i++ {
		d = append(d, p.RollDice())
	}

	return d
}

func Play(players_count, dices_count int) {
	if players_count < 2 || dices_count < 1 {
		log.Fatal("Invalid number of players or dices")
	}

	players = make([]Player, players_count)

	for {
		for i := 0; i < players_count; i++ {
			// Create players
			players[i].SetNumber(i + 1)

			// Roll dices
			if dices := len(players[i].Dice); dices == 0 {
				players[i].Dice = players[i].RollDices(dices_count)
			} else {
				dices_count = 0
				players[i].Dice = players[i].RollDices(dices)
			}
		}

		// Print result
		fmt.Println(`=== ROLL DICES ===`)
		for i := 0; i < players_count; i++ {
			fmt.Printf("Player %d: %v\n", players[i].Number, players[i].Dice)
		}

		// Evaluate
		fmt.Printf("\n=== EVALUATE ===\n")
		var tmp_len []int

		// Calculate score if there is available number 6 in dices
		for i := 0; i < players_count; i++ {
			for j := 0; j < len(players[i].Dice); j++ {
				if players[i].Dice[j] == 6 {
					players[i].Score++

					// Remove number 6 from dices
					players[i].Dice = append(players[i].Dice[:j], players[i].Dice[j+1:]...)
					j--
				}

			}

			tmp_len = append(tmp_len, len(players[i].Dice))
		}

		// Send number 1 to the next player and remove it from current player
		for i := 0; i < players_count; i++ {
			var counter int = 0
			for j := 0; j < len(players[i].Dice); j++ {
				counter++
				if tmp_len[i] >= counter {
					if players[i].Dice[j] == 1 {
						// Distribute 1 to the next player and remove it from current player
						nextPlayerIndex := (i + 1) % players_count
						players[nextPlayerIndex].Dice = append(players[nextPlayerIndex].Dice, 1)

						// Remove number 1 from dices
						players[i].Dice = append(players[i].Dice[:j], players[i].Dice[j+1:]...)
						j--
					}
				}
			}
		}

		for i := 0; i < players_count; i++ {
			fmt.Printf("Player %d (%d): %v\n", players[i].Number, players[i].Score, players[i].Dice)
		}

		// Determine winner if there is only one player left
		var active_player int

		for i := 0; i < players_count; i++ {
			if len(players[i].Dice) == 0 {
				active_player++
			}
		}

		if players_count-active_player == 1 {
			fmt.Printf("\n=== WINNER ===\n")
			// Find Max score
			var max_score int
			for i := 0; i < players_count; i++ {
				if players[i].Score > max_score {
					max_score = players[i].Score
				}
			}

			// Find winner
			for i := 0; i < players_count; i++ {
				if players[i].Score == max_score {
					fmt.Printf("Player %d (%d)\n", players[i].Number, players[i].Score)
				}
			}

			break
		}

		fmt.Printf("\n=== NEXT ROUND ===\n\n")

		// Delay 1 second
		time.Sleep(1 * time.Second)
	}
}

func main() {
	Play(3, 5)
}
