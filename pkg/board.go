package pkg

import (
	"fmt"
)

// has a matrix
type Board struct {
	Grid [3][3]PlayerSymbol
	//MoveCount int
}

func NewBoard() *Board {

	board := &Board{}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			board.Grid[i][j] = Empty
		}
	}
	return board

}

func (b *Board) PrintBoard() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print(string(b.Grid[i][j]), " | ")
		}
		fmt.Print("\n")
	}
}
func (b *Board) IsFull() bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if b.Grid[i][j] == Empty {
				return false

			}
		}

	}
	return true
}

func (b *Board) CheckWin() bool {
	for i := 0; i < 3; i++ {
		if b.Grid[i][0] != Empty && b.Grid[i][0] == b.Grid[i][1] && b.Grid[i][1] == b.Grid[i][2] {
			return true
		}
	}

	for i := 0; i < 3; i++ {
		if b.Grid[0][i] != Empty && b.Grid[0][i] == b.Grid[1][i] && b.Grid[1][i] == b.Grid[2][i] {
			return true
		}
	}

	if b.Grid[0][0] != Empty && b.Grid[0][0] == b.Grid[1][1] && b.Grid[1][1] == b.Grid[2][2] {
		return true
	}
	return b.Grid[2][0] != Empty && b.Grid[2][0] == b.Grid[1][1] && b.Grid[1][1] == b.Grid[0][2]
}

func (b *Board) IsValidMove(x, y int) bool {
	if x < 0 || x > 2 || y < 0 || y > 2 || b.Grid[x][y] != Empty {
		return false
	}
	return true
}

func (b *Board) UpdateBoard(x, y int, p PlayerSymbol) {

	b.Grid[x][y] = p

}
