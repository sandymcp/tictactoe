package tictactoe

import "errors"

type Cell byte

const (
	E    Cell = '.'
	X    Cell = 'X'
	O    Cell = '0'
	side      = 3
)

// assumes a 3*3 matrix
var wins = [][]int{
	{0, 1, 2},
	{3, 4, 5},
	{6, 7, 8},
	{0, 3, 6},
	{1, 4, 7},
	{2, 5, 8},
	{0, 4, 8},
	{2, 4, 6},
}

type Board [side * side]Cell

func NewBoard() Board {
	return Board{E, E, E, E, E, E, E, E, E}
}

var WrongMove = errors.New("this move has already been played")
var XExceeded = errors.New("the X coordinate is too big")
var YExceeded = errors.New("the Y coordinate is too big")

// Move assumes a valid move is passed
// Checks if the cell requested by the move is
func (b *Board) Move(m Move) error {
	x, y := m.ToCoords()
	if (0 > x) || (x > side) {
		return XExceeded
	}
	if (0 > y) || (y > side) {
		return YExceeded
	}
	index := x*side + y
	switch b[index] {
	case E:
		b[index] = m.ToCell()
	default:
		return WrongMove
	}
	return nil
}

func (b Board) CheckWinner(p Player) bool {
	for _, win := range wins {
		winner := true
		for _, cell := range win {
			if p.ToCell() != b[cell] {
				winner = false
				break
			}
		}
		if winner {
			return true
		}
	}
	return false
}

func (b Board) CheckDraw() bool {
	for _, cell := range b {
		if cell == E {
			return false
		}
	}
	return true
}
