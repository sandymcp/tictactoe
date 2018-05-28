package tictactoe

import "errors"

// XCoord is a letter (like chess notation)
type XCoord byte

// Ycoord is a number
type YCoord uint8

const (
	X1 XCoord = 'a'
	X2 XCoord = 'b'
	X3 XCoord = 'c'
	Y1 YCoord = '1'
	Y2 YCoord = '2'
	Y3 YCoord = '3'
)

// Player a player either plays as X or O
type Player Cell

const (
	PX = Player(X)
	PO = Player(O)
)

type Move struct {
	Player
	XCoord
	YCoord
}

func (p Player) ToCell() Cell {
	return Cell(p)
}

// ToCoords gives the cartesian coordinates of where to place the move on the board
func (m *Move) ToCoords() (x, y int) {
	return int(m.XCoord - X1), int(m.YCoord - Y1)
}

var InvalidPlayer = errors.New("invalid player: valids are 1|2")
var InvalidXCoord = errors.New("invalid X: valid are A|B|C")
var InvalidYCoord = errors.New("invalid Y: valid are 1|2|3")

func (m *Move) IsValid() error {
	if err := m.Player.IsValid(); err != nil {
		return err
	}
	if err := m.XCoord.IsValid(); err != nil {
		return err
	}
	if err := m.YCoord.IsValid(); err != nil {
		return err
	}
	return nil
}

func (p Player) IsValid() error {
	switch p {
	case PX, PO:
		return nil
	default:
		return InvalidPlayer
	}
}

func (x XCoord) IsValid() error {
	switch x {
	case X1, X2, X3:
		return nil
	default:
		return InvalidXCoord
	}
}

func (y YCoord) IsValid() error {
	switch y {
	case Y1, Y2, Y3:
		return nil
	default:
		return InvalidYCoord
	}
}
