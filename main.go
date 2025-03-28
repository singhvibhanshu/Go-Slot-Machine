package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// checkWin determines winnings for each row based on matching symbols.
func checkWin(spin [][]string, multipliers map[string]uint) []uint {
	lines := []uint{}

	for _, row := range spin {
		win := true
		checkSymbol := row[0]
		for _, symbol := range row[1:] {
			if checkSymbol != symbol {
				win = false
				break
			}
		}
		if win {
			lines = append(lines, multipliers[checkSymbol])
		} else {
			lines = append(lines, 0)
		}
	}
	return lines
}

// GetName reads the player's name from stdin using bufio for safety.
func GetName() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to Tim's Casino Slot Machine!")
	fmt.Print("Enter your name: ")
	name, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading name. Using default name 'Player'.")
		return "Player"
	}
	name = strings.TrimSpace(name)
	if name == "" {
		name = "Player"
	}
	fmt.Printf("Welcome %s, let's play!\n", name)
	return name
}

// GetBet prompts the player to enter a bet amount, ensuring valid input.
func GetBet(balance uint) uint {
	for {
		fmt.Printf("Enter your bet (balance = $%d) or 0 to quit: ", balance)
		var bet uint
		_, err := fmt.Scanln(&bet)
		if err != nil {
			// Clear out invalid input
			fmt.Println("Please enter a valid number.")
			// Discard the remainder of the line to avoid looping
			var discard string
			fmt.Scanln(&discard)
			continue
		}

		if bet > balance {
			fmt.Println("Bet cannot be larger than your current balance.")
		} else {
			return bet
		}
	}
}

// GenerateSymbolArray creates a symbol array based on the given counts.
func GenerateSymbolArray(symbols map[string]uint) []string {
	symbolArr := []string{}
	for symbol, count := range symbols {
		for i := uint(0); i < count; i++ {
			symbolArr = append(symbolArr, symbol)
		}
	}
	return symbolArr
}

// getRandomNumber generates a random number in the given range.
func getRandomNumber(min int, max int) int {
	return rand.Intn(max-min+1) + min
}

// GetSpin creates a 2D slice representing the slot machine spin.
func GetSpin(reel []string, rows int, cols int) [][]string {
	result := make([][]string, rows)
	for i := 0; i < rows; i++ {
		result[i] = []string{}
	}

	for col := 0; col < cols; col++ {
		selected := map[int]bool{}
		for row := 0; row < rows; row++ {
			for {
				randomIndex := getRandomNumber(0, len(reel)-1)
				if !selected[randomIndex] {
					selected[randomIndex] = true
					result[row] = append(result[row], reel[randomIndex])
					break
				}
			}
		}
	}
	return result
}

// PrintSpin prints the current spin in a formatted way.
func PrintSpin(spin [][]string) {
	fmt.Println("----------")
	for _, row := range spin {
		for j, symbol := range row {
			fmt.Print(symbol)
			if j != len(row)-1 {
				fmt.Print(" | ")
			}
		}
		fmt.Println()
	}
	fmt.Println("----------")
}

func main() {
	// Seed the random number generator.
	rand.Seed(time.Now().UnixNano())

	// Define symbols and multipliers.
	symbols := map[string]uint{
		"A": 4,
		"B": 7,
		"C": 12,
		"D": 20,
	}
	multipliers := map[string]uint{
		"A": 20,
		"B": 10,
		"C": 5,
		"D": 2,
	}

	// Generate the reel array.
	symbolArr := GenerateSymbolArray(symbols)

	// Set initial balance.
	balance := uint(200)

	// Welcome the player.
	GetName()

	// Game loop.
	for balance > 0 {
		bet := GetBet(balance)
		if bet == 0 {
			fmt.Println("Thanks for playing!")
			break
		}

		balance -= bet

		// Show the spin result.
		spin := GetSpin(symbolArr, 3, 3)
		PrintSpin(spin)

		// Calculate winnings.
		winningLines := checkWin(spin, multipliers)
		totalWin := uint(0)
		for i, multi := range winningLines {
			win := multi * bet
			totalWin += win
			if multi > 0 {
				fmt.Printf("Line #%d wins: $%d (%dx multiplier)\n", i+1, win, multi)
			}
		}
		if totalWin == 0 {
			fmt.Println("No winning lines this spin. Better luck next time!")
		} else {
			balance += totalWin
		}

		fmt.Printf("New balance: $%d\n\n", balance)
	}

	fmt.Printf("Game over! You left with $%d.\n", balance)
}
