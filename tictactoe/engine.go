package tictactoe

import "github.com/pkg/errors"

type State struct {
	Board
	NextPlayer Player
}

func NewState() *State{
	return &State{Board:NewBoard(),NextPlayer:PX}
}

var WrongPlayer = errors.New("Other player's turn")

func (state *State) Play(move Move) (win, draw bool, err error) {
	if err = move.IsValid(); err != nil {
		return false,false, err
	}
	if move.Player != state.NextPlayer {
		return false, false, WrongPlayer
	}
	if err := state.Board.Move(move); err != nil {
		return false, false, err
	}
	if state.Board.CheckWinner(move.Player) {
		return true, false, nil
	}
	if state.Board.CheckDraw() {
		return false, true, nil
	}
	switch move.Player {
	case PX:
		state.NextPlayer = PO
	case PO:
		state.NextPlayer = PX
	default:
		panic("Invalid player, validation must have failed...")
	}
	return false, false, nil
}
