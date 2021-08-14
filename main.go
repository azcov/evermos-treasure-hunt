package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

const (
	gridMaxX   = 8
	gridMaxY   = 6
	OpeningMsg = `  // ---------------------- //
 // Treasure Hunt by Mufid //
// ---------------------- //
How to Play :
- Find the treasure ($)
- Press "w" to move player (X) up
- Press "s" to move player (X) down
- Press "d" to move player (X) right
- Press "r" to reset the game
Play now? (y/n):`
	TextColorRed    = "\033[31m%s\033[0m"
	TextColorYellow = "\033[33m%s\033[0m"
	TextColorBlue   = "\033[34m%s\033[0m"
	TextColorWhite  = "\033[97m%s\033[0m"
)

func clearConsole() {
	switch runtime.GOOS {
	case "linux":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		fmt.Println("Cannot Clear Console")
	}
}

type TreasureHunt struct {
	grid     Grid
	player   Position
	treasure Position
	msg      string
	isFind   bool
}

func NewGame() TreasureHunt {
	return TreasureHunt{
		grid: [][]bool{
			{false, false, false, false, false, false, false, false},
			{false, true, true, true, true, true, true, false},
			{false, true, false, false, false, true, true, false},
			{false, true, true, true, false, true, false, false},
			{false, true, false, true, true, true, true, false},
			{false, false, false, false, false, false, false, false},
		},
		player:   NewPlayer(),
		treasure: NewTreasure(),
		msg:      "Lets Find Treasure!",
	}
}
func (t *TreasureHunt) clearMsg() {
	t.msg = ""
}

func (t *TreasureHunt) PrintGrid() {
	clearConsole()
	fmt.Println(t.msg)
	for y := range t.grid {
		for x := range t.grid[y] {
			if t.player.x == x && t.player.y == y {
				fmt.Printf(TextColorBlue, "X ")
				continue
			}
			if t.treasure.x == x && t.treasure.y == y {
				fmt.Printf(TextColorYellow, "$ ")
				continue
			}
			if t.grid[y][x] {
				fmt.Printf(TextColorWhite, ". ")
			} else {
				fmt.Printf(TextColorRed, "# ")
			}
		}
		fmt.Println()
	}
}

func (t *TreasureHunt) PlayerUp() {
	newPos := Position{t.player.x, t.player.y - 1}
	if !t.grid[newPos.y][newPos.x] {
		t.msg = "cannot move up"
	} else {
		t.player = newPos
		t.clearMsg()
		t.isFindTreasure()
	}
}

func (t *TreasureHunt) PlayerDown() {
	newPos := Position{t.player.x, t.player.y + 1}
	if !t.grid[newPos.y][newPos.x] {
		t.msg = "cannot move down"
	} else {
		t.player = newPos
		t.clearMsg()
		t.isFindTreasure()
	}
}

func (t *TreasureHunt) PlayerRight() {
	newPos := Position{t.player.x + 1, t.player.y}
	if !t.grid[newPos.y][newPos.x] {
		t.msg = "cannot move right"
	} else {
		t.player = newPos
		t.clearMsg()
		t.isFindTreasure()
	}
}

func (t *TreasureHunt) isFindTreasure() {
	if t.player.x == t.treasure.x && t.player.y == t.treasure.y {
		t.msg = "Yeppy! you find the Treasure!"
		t.isFind = true
	}
}

type Grid [][]bool

type Position struct {
	x int
	y int
}

func NewPlayer() Position {
	return Position{x: 1, y: 4}
}
func NewTreasure() Position {
	return Position{x: 6, y: 4}
}

func main() {
	clearConsole()
	fmt.Printf(TextColorYellow, OpeningMsg)
	reader := bufio.NewReader(os.Stdin)
	cmd, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	// convert CRLF to LF
	cmd = strings.Replace(cmd, "\n", "", -1)

	switch strings.ToLower(cmd) {
	case "y":
		treasureHunt := NewGame()
		for {
			treasureHunt.PrintGrid()

			if treasureHunt.isFind {
				fmt.Println("Game End")
				break
			}

			cmd, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println(err)
				break
			}
			// convert CRLF to LF
			cmd = strings.Replace(cmd, "\n", "", -1)

			switch strings.ToLower(cmd) {
			case "w":
				treasureHunt.PlayerUp()
			case "d":
				treasureHunt.PlayerRight()
			case "s":
				treasureHunt.PlayerDown()
			case "r":
				treasureHunt = NewGame()
			default:
				fmt.Println(cmd)
				time.Sleep(time.Duration(1 * time.Second))
			}
		}
	default:
		break
	}
	fmt.Println("Bye bye!")
}
