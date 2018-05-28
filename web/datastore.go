package web

import "github.com/sandymcp/tictactoe/tictactoe"

type Datastore interface{
	NewID() string
	Upsert(gameID string, state *tictactoe.State)
	Get(gameID string, ) *tictactoe.State
	Delete(gameID string)
}