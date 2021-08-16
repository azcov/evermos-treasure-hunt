# Treasure Hunt

Build a simple command-line program for helping the user hunt for a treasure that satisfies the following requirements:
1. Build a simple grid with the following layout:
```sh
########
#......#
#.###..#
#...#.##
#X#....#
########
```
- `#` represents an obstacle.
- . represents a clear path.
- X represents the player’s starting position.
- A treasure is hidden within one of the clear path points, and the user must find it.
From the starting position, the user must navigate in a specific order:
- Up/North A step(s), then
- Right/East B step(s), then
- Down/South C step(s).
The program must output a list of probable coordinate points where the treasure might be located.
Bonus points: display the grid with all the probable treasure locations marked with a $ symbol.
## How To Run

Make sure you have installed Golang, If you haven't installed golang please [install](https://golang.org/doc/install) first.

If you have installed Make (software):

```sh
$ make play
```

If you haven't installed Make (software):

```sh
$ go run main.go
```



## How To Play

How to Play :
- Find the treasure ($)
- Input a number (n) to move player (X) up 
- Input a number (n) to move player(X) down
- Input a number (n) to move player (X) right
