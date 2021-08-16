package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

const (
	gridMaxX   = 8
	gridMaxY   = 6
	OpeningMsg = `  // ---------------------- //
 // Treasure Hunt by Mufid //
// ---------------------- //
How to Play :
- Find the treasure ($)
- Input a number (n) to move player (X) up 
- Input a number (n) to move player (X) down
- Input a number (n) to move player (X) right

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
	treasure []Position
	msg      string
	isFind   bool
	step     int
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
		step:     1,
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
			isTreasurePosition := false
			if t.player.x == x && t.player.y == y {
				fmt.Printf(TextColorBlue, "X ")
				continue
			}
			for z := range t.treasure {
				if t.treasure[z].x == x && t.treasure[z].y == y {
					fmt.Printf(TextColorYellow, "$ ")
					isTreasurePosition = true
					break
				}
			}
			if !isTreasurePosition {
				if t.grid[y][x] {
					fmt.Printf(TextColorWhite, ". ")
				} else {
					fmt.Printf(TextColorRed, "# ")
				}
			}

		}
		fmt.Println()
	}
	switch t.step {
	case 1:
		fmt.Print("input step for move north/up   : ")
	case 2:
		fmt.Print("input step for move right/east : ")
	case 3:
		fmt.Print("input step for move down/south : ")
	}
}

func (t *TreasureHunt) InputStep(cmd string) {
	if strings.ToLower(cmd) != "r" {
		steps, err := strconv.Atoi(cmd)
		if err != nil {
			t.msg = "input not allowed"
		}
		if steps > 0 {
			switch t.step {
			case 1:
				t.PlayerUp(steps)
				t.step++
			case 2:
				t.PlayerRight(steps)
				t.step++
			case 3:
				t.PlayerDown(steps)
				t.step++
			}
		} else {
			t.msg = "input must be greater than 0"
		}
	} else {
		*t = NewGame()
	}
}

func (t *TreasureHunt) PlayerUp(n int) {
	for i := 0; i < n; i++ {
		newPos := Position{t.player.x, t.player.y - 1}
		if !t.grid[newPos.y][newPos.x] {
			t.msg = "cannot move up again"
			break
		} else {
			t.player = newPos
			t.clearMsg()
			t.isFindTreasure()
		}
	}
}

func (t *TreasureHunt) PlayerDown(n int) {
	for i := 0; i < n; i++ {
		newPos := Position{t.player.x, t.player.y + 1}
		if !t.grid[newPos.y][newPos.x] {
			t.msg = "cannot move down"
		} else {
			t.player = newPos
			t.clearMsg()
			t.isFindTreasure()
		}
	}
}

func (t *TreasureHunt) PlayerRight(n int) {
	for i := 0; i < n; i++ {
		newPos := Position{t.player.x + 1, t.player.y}
		if !t.grid[newPos.y][newPos.x] {
			t.msg = "cannot move right"
		} else {
			t.player = newPos
			t.clearMsg()
			t.isFindTreasure()
		}
	}
}

func (t *TreasureHunt) isFindTreasure() {
	for i := range t.treasure {
		if t.player.x == t.treasure[i].x && t.player.y == t.treasure[i].y {
			t.msg = "Yeppy! you find the Treasure!"
			t.isFind = true
		}
	}
	if !t.isFind && t.step == 4 {
		t.msg = "you failed to find the Treasure!"
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
func NewTreasure() []Position {
	return []Position{
		{x: 3, y: 4},
		{x: 5, y: 4},
		{x: 6, y: 2},
	}
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

			if treasureHunt.isFind || treasureHunt.step == 4 {
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
			treasureHunt.InputStep(cmd)
		}
	default:
		break
	}
	fmt.Println("Bye bye!")
}
