package tictactoe

import (
	"testing"
	"math"
)

func TestPlayer_IsValid(t *testing.T) {
	for p := Player(0); p < 255; p++ {
		err := p.IsValid()
		switch p {
		case PX, PO:
			if got, want := err, error(nil); got != want {
				t.Errorf("Invalid player: got %v,want %v", got, want)
			}
		default:
			if got, want := err, InvalidPlayer; got != want {
				t.Errorf("Invalid player: got %v,want %v", got, want)
			}
		}
	}
}

func TestXCoord_IsValid(t *testing.T) {
	for p := XCoord(0); p < 255; p++ {
		err := p.IsValid()
		switch p {
		case X1, X2, X3:
			if got, want := err, error(nil); got != want {
				t.Errorf("Invalid x-coord: got %v,want %v", got, want)
			}
		default:
			if got, want := err, InvalidXCoord; got != want {
				t.Errorf("Invalid y-coord: got %v,want %v", got, want)
			}
		}
	}
}

func TestYCoord_IsValid(t *testing.T) {
	for p := YCoord(0); p < 255; p++ {
		err := p.IsValid()
		switch p {
		case Y1, Y2, Y3:
			if got, want := err, error(nil); got != want {
				t.Errorf("Invalid x-coord: got %v,want %v", got, want)
			}
		default:
			if got, want := err, InvalidYCoord; got != want {
				t.Errorf("Invalid y-coord: got %v,want %v", got, want)
			}
		}
	}
}

func TestMove_IsValid(t *testing.T) {
	tests := map[string]struct {
		Move
		wantError error
	}{
		"no error": {
			Move: Move{Player: PO, XCoord: X1, YCoord: Y1},
		},
		"invalid player": {
			Move:      Move{Player: Player(math.MaxUint8), XCoord: X2, YCoord: Y2},
			wantError: InvalidPlayer,
		},
		"invalid x-coord": {
			Move:      Move{Player: PX, XCoord: XCoord(math.MaxUint8), YCoord: Y3},
			wantError: InvalidXCoord,
		},
		"invalid y-coord": {
			Move:      Move{Player: PX, XCoord: X3, YCoord: YCoord(math.MaxUint8)},
			wantError: InvalidYCoord,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			err := test.Move.IsValid()
			if got, want := err, test.wantError; got != want {
				t.Errorf("Invalid move validation: got %v,want %v", got, want)
			}
		})
	}
}

func TestMove_ToCoords(t *testing.T) {
	tests := map[string]struct {
		Move
		x int
		y int
	}{
		"A1": {Move{PO, X1, Y1}, 0,0},
		"A2": {Move{PO, X1, Y2}, 0,1},
		"A3": {Move{PO, X1, Y3}, 0,2},
		"B1": {Move{PO, X2, Y1}, 1,0},
		"B2": {Move{PO, X2, Y2}, 1,1},
		"B3": {Move{PO, X2, Y3}, 1,2},
		"B4": {Move{PO, X3, Y1}, 2,0},
		"B5": {Move{PO, X3, Y2}, 2,1},
		"B6": {Move{PO, X3, Y3}, 2,2},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			x, y := test.ToCoords()
			if got, want := x, test.x;  got != want {
				t.Errorf("Invalid x-coord: got %v,want %v", got, want)
			}
			if got, want := y, test.y;  got != want {
				t.Errorf("Invalid y-coord: got %v,want %v", got, want)
			}
		})
	}
}
