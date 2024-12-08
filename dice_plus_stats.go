package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"time"
)

// ClearTerminal clears the terminal screen
func ClearTerminal() {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

// GetDiceFace returns ASCII art for a dice face
func GetDiceFace(face int) string {
	diceFaces := map[int]string{
		1: `
+-------+
|       |
|   *   |
|       |
+-------+`,
		2: `
+-------+
| *     |
|       |
|     * |
+-------+`,
		3: `
+-------+
| *     |
|   *   |
|     * |
+-------+`,
		4: `
+-------+
| *   * |
|       |
| *   * |
+-------+`,
		5: `
+-------+
| *   * |
|   *   |
| *   * |
+-------+`,
		6: `
+-------+
| *   * |
| *   * |
| *   * |
+-------+`,
	}
	return diceFaces[face]
}

// DisplayHistory prints the last N rolls in a single line
func DisplayHistory(history []int) {
	fmt.Print("\nLast 10 Rolls: ")
	for _, roll := range history {
		fmt.Printf("%d ", roll)
	}
	fmt.Println()
}

// CheckRareEvent checks for rare events like rolling five 6's in a row
func CheckRareEvent(history []int) {
	if len(history) < 5 {
		return
	}
	isRareEvent := true
	for i := len(history) - 5; i < len(history); i++ {
		if history[i] != 6 {
			isRareEvent = false
			break
		}
	}
	if isRareEvent {
		fmt.Println("\n🎉🎉 RARE EVENT: Five 6's in a row! 🎉🎉")
		fmt.Println("Congratulations on this amazing streak!")
		// Display the probability of the rare event
		probability := math.Pow(1.0/6.0, 5) * 100
		fmt.Printf("Probability of this event: %.8f%%\n", probability)
	}
}

// SimulateDiceThrows simulates dice rolls and updates stats
func SimulateDiceThrows() {
	// Initialize stats
	diceRollStats := make(map[int]int)
	var history []int
	totalRolls := 0
	historyLimit := 10 // Limit of the rolling history

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Start the simulation
	for {
		// Simulate a dice roll (1 to 6)
		roll := rand.Intn(6) + 1
		diceRollStats[roll]++
		totalRolls++

		// Update history
		history = append(history, roll)
		if len(history) > historyLimit {
			history = history[1:] // Remove oldest roll to maintain the limit
		}

		// Clear terminal for updated stats
		ClearTerminal()

		// Display the roll and ASCII art
		fmt.Printf("***** Rolled a %d *****", roll)
		fmt.Println(GetDiceFace(roll))
		fmt.Println("\nDice Roll Stats:")
		fmt.Printf("%-10s%-10s%-10s\n", "Dice Face", "Count", "Percentage")
		fmt.Println("===============================")

		for face := 1; face <= 6; face++ {
			count := diceRollStats[face]
			percentage := 0.0
			if totalRolls > 0 {
				percentage = (float64(count) / float64(totalRolls)) * 100
			}
			fmt.Printf("%-10d%-10d%-10.2f%%\n", face, count, percentage)
		}

		// Display rolling history
		DisplayHistory(history)

		// Check for rare events
		CheckRareEvent(history)
		fmt.Println("")
		fmt.Printf("Total Rolls: %d\n", totalRolls)

		// Wait for 2 seconds
		time.Sleep(2 * time.Second)
	}
}

func main() {
	SimulateDiceThrows()
}
