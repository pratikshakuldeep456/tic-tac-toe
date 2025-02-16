package main

import "pratikshakuldeep456/tic-tac-toe/pkg"

func main() {

	p1 := pkg.NewPlayer("pratiksha", pkg.X)
	p2 := pkg.NewPlayer("ankit", pkg.O)
	game := pkg.NewGame(p1, p2)

	game.StartGame()

}
