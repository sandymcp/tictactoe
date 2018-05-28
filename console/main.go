package main

import (
	"fmt"

	"github.com/sandymcp/tictactoe/tictactoe"
)

func main() {
	fmt.Printf("Welcome to Tictatoe, player X begins\n")
	s := tictactoe.NewState()
	fmt.Printf("Enter a move valid moves are a1,a2,a3,b1,b2,b3,c1,c2,c3\n")
	for  {
		paint(s.Board)
		m := tictactoe.Move{Player: s.NextPlayer}
		fmt.Printf("Player %c move please: ", m.Player)
		fmt.Scanf("%c%c\n", &m.XCoord, &m.YCoord)
		fmt.Printf("\n")
		win, draw, err := s.Play(m)
		if err != nil {
			fmt.Printf("%v\n", err)
			continue
		}
		if win {
			fmt.Printf("Congratulations %c you have won the game\n", m.Player)
			break
		}
		if draw {
			fmt.Printf("Hard luck a draw\n")
		}
	}

}

func paint(b tictactoe.Board) {
	fmt.Printf(" %c%c%c", tictactoe.X1, tictactoe.X2, tictactoe.X3)
	yLabels := []tictactoe.YCoord{tictactoe.Y1, tictactoe.Y2, tictactoe.Y3}
	for i, c := range b {
		if i%3 == 0 {
			fmt.Printf("\n%c", yLabels[i/3])
		}
		fmt.Printf("%c", c)
	}
	fmt.Printf("\n")
}
