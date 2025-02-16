package pkg

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	Board         *Board
	Players       []Player
	CurrentPlayer Player
}

func NewGame(p1, p2 Player) *Game {
	return &Game{
		Board:         NewBoard(),
		Players:       []Player{p1, p2},
		CurrentPlayer: p1,
	}

}

func (g *Game) StartGame() {

	g.Board.PrintBoard()

	for !g.Board.IsFull() && !g.Board.CheckWin() {
		g.Board.PrintBoard()
		fmt.Println("current player is", g.CurrentPlayer.Name)

		x, y := g.InputMove()
		if g.Board.IsValidMove(x, y) {
			g.Board.UpdateBoard(x, y, g.CurrentPlayer.Symbol)
		} else {
			fmt.Println("Invalid move")
			continue
		}

		g.SwitchPlayer()
		if g.Board.CheckWin() {
			fmt.Println("winner is")
		}

	}

	if g.Board.CheckWin() {
		fmt.Println("winner is", g.CurrentPlayer.Name)
	} else {
		fmt.Println("Game is a draw")
	}

}

func (g *Game) SwitchPlayer() {
	if g.CurrentPlayer == g.Players[0] {
		g.CurrentPlayer = g.Players[1]
	} else {
		g.CurrentPlayer = g.Players[0]
	}
}

func (g *Game) InputMove() (int, int) {

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("ENTER YOUR MOVE x,y in b/w range 0-2")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		coords := strings.Split(input, ",")

		if len(coords) != 2 {
			fmt.Println("Invalid input")
			continue
		}

		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])

		if x < 0 || x > 2 || y < 0 || y > 2 {
			fmt.Println("Invalid input")
			continue
		}

		return x, y

	}

}
