package tictactoe

import (
	"testing"
	"reflect"
)

func TestBoard_Move(t *testing.T) {
	tests := map[string]struct {
		moves  []Move
		errs   []error
		result Board
	}{
		"Simple Move 1": {
			moves:  []Move{{PX, X1, Y1}},
			errs:   []error{nil},
			result: Board{X, E, E, E, E, E, E, E, E},
		},
		"Simple Move 2": {
			moves:  []Move{{PX, X1, Y2}},
			errs:   []error{nil},
			result: Board{E, X, E, E, E, E, E, E, E},
		},
		"Simple Move 3": {
			moves:  []Move{{PX, X1, Y3}},
			errs:   []error{nil},
			result: Board{E, E, X, E, E, E, E, E, E},
		},
		"Simple Move 4": {
			moves:  []Move{{PX, X2, Y1}},
			errs:   []error{nil},
			result: Board{E, E, E, X, E, E, E, E, E},
		},
		"Simple Move 5": {
			moves:  []Move{{PX, X2, Y2}},
			errs:   []error{nil},
			result: Board{E, E, E, E, X, E, E, E, E},
		},
		"Simple Move 6": {
			moves:  []Move{{PX, X2, Y3}},
			errs:   []error{nil},
			result: Board{E, E, E, E, E, X, E, E, E},
		},
		"Simple Move 7": {
			moves:  []Move{{PX, X3, Y1}},
			errs:   []error{nil},
			result: Board{E, E, E, E, E, E, X, E, E},
		},
		"Simple Move 8": {
			moves:  []Move{{PX, X3, Y2}},
			errs:   []error{nil},
			result: Board{E, E, E, E, E, E, E, X, E},
		},
		"Simple Move 9": {
			moves:  []Move{{PX, X3, Y3}},
			errs:   []error{nil},
			result: Board{E, E, E, E, E, E, E, E, X},
		},
		"Two Good Moves": {
			moves:  []Move{{PX, X3, Y3}, {PO, X2, Y2}},
			errs:   []error{nil, nil},
			result: Board{E, E, E, E, O, E, E, E, X},
		},
		"Other Good Moves": {
			moves: []Move{
				{PX, X1, Y3}, {PO, X1, Y2}, {PX, X3, Y1},
			},
			errs:   []error{nil, nil, nil},
			result: Board{E, O, X, E, E, E, X, E, E},
		},
		"Repeat Move XO": {
			moves:  []Move{{PX, X2, Y2}, {PO, X2, Y2}},
			errs:   []error{nil, WrongMove},
			result: Board{E, E, E, E, X, E, E, E, E},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			b := NewBoard()
			errs := make([]error, len(test.errs))
			for i, move := range test.moves {
				errs[i] = b.Move(move)
			}
			if got, want := errs, test.errs; !reflect.DeepEqual(got, want) {
				t.Errorf("Invalid errors:\ngot.. %v\nwant. %v", got, want)
			}
			if got, want := b, test.result; !reflect.DeepEqual(got, want) {
				t.Errorf("Invalid board:\ngot.. %v\nwant. %v", got, want)
			}
		})
	}
}

func TestBoard_CheckWinner(t *testing.T) {
	tests := map[string]struct {
		board  Board
		player Player
		win    bool
	}{
		"easy win X d 1":  {board: Board{X, O, E, E, X, E, E, O, X}, player: PX, win: true},
		"easy win X h 1":  {board: Board{X, X, X, E, O, E, E, O, O}, player: PX, win: true},
		"easy win O h 2":  {board: Board{X, O, E, O, O, O, E, E, X}, player: PO, win: true},
		"easy win X h 3":  {board: Board{X, O, E, E, X, E, X, X, X}, player: PX, win: true},
		"easy win X d 2":  {board: Board{E, O, X, E, X, E, X, O, E}, player: PX, win: true},
		"easy win O v 1":  {board: Board{O, X, E, O, X, E, O, E, X}, player: PO, win: true},
		"easy win X v 2":  {board: Board{O, X, E, O, X, E, O, X, O}, player: PX, win: true},
		"easy win O v 3":  {board: Board{X, E, O, E, X, O, X, X, O}, player: PO, win: true},
		"not quite X d 1": {board: Board{E, O, E, E, X, E, E, O, X}, player: PX, win: false},
		"not quite X h 1": {board: Board{X, E, X, E, O, E, E, O, O}, player: PX, win: false},
		"not quite O h 2": {board: Board{X, O, E, O, X, O, E, E, X}, player: PO, win: false},
		"not quite X h 3": {board: Board{O, O, E, E, X, E, X, E, X}, player: PX, win: false},
		"not quite X d 2": {board: Board{E, O, X, E, O, E, X, O, E}, player: PX, win: false},
		"not quite O v 1": {board: Board{O, X, E, O, X, E, E, E, X}, player: PO, win: false},
		"not quite X v 2": {board: Board{O, X, E, O, E, E, O, X, O}, player: PX, win: false},
		"not quite O v 3": {board: Board{X, E, O, E, X, E, X, X, O}, player: PO, win: false},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			win := test.board.CheckWinner(test.player)
			if got, want := win, test.win; got != want {
				t.Errorf("Invalid win: got %v, want %v", got, want)
			}
		})
	}
}
