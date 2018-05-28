package tictactoe

import (
	"testing"
	"reflect"
)

func TestPlay(t *testing.T) {
	tests := map[string]struct {
		state State
		move  Move
		err   error
		win   bool
		draw   bool
	}{
		"Winning Move": {
			state:
			State{Board: Board{E, X, X, E, E, E, E, E, E}, NextPlayer: PX},
			move: Move{Player: PX, XCoord: X1, YCoord: Y1},
			win:  true,
		},
		"Draw ": {
			state:
			State{Board: Board{X, X, O, O, O, X, X, O, E}, NextPlayer: PX},
			move: Move{Player: PX, XCoord: X3, YCoord: Y3},
			win:  false,
			draw: true,
		},
		"Wrong Player": {
			state:
			State{Board: Board{E, X, X, E, E, E, E, E, E}, NextPlayer: PX},
			move: Move{Player: PO, XCoord: X1, YCoord: Y1},
			win:  false,
			err: WrongPlayer,
		},
		"Bad Move": {
			state:
			State{Board: Board{X, X, E, E, E, E, E, E, E}, NextPlayer: PX},
			move: Move{Player: PX, XCoord: X1, YCoord: Y2},
			win:  false,
			err: WrongMove,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			win, draw, err := test.state.Play(test.move)
			if got, want := err, test.err; !reflect.DeepEqual(got, want) {
				t.Errorf("Invalid error: got %v, want %v", got, want)
			}
			if got, want := win, test.win; got != want {
				t.Errorf("Invalid win: got %v, want %v", got, want)
			}
			if got, want := draw, test.draw; got != want {
				t.Errorf("Invalid draw: got %v, want %v", got, want)
			}
		})
	}
}
