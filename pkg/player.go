package pkg

type PlayerSymbol rune

const (
	Empty PlayerSymbol = ' '
	X     PlayerSymbol = 'X'
	O     PlayerSymbol = 'O'
)

type Player struct {
	Name   string
	Symbol PlayerSymbol
}

func NewPlayer(name string, symbol PlayerSymbol) Player {
	return Player{Name: name, Symbol: symbol}
}
