package web

import "github.com/sandymcp/tictactoe/tictactoe"

type Notifier interface{
	Register(player tictactoe.Player)
	// Push notifies the client that the state has changed, assumes the client knows which player it represents
	Push(gameid string, board tictactoe.State)
	// Listen for moves on either a REST API or a WebSocket or such like, retrieves from the data store and
	// plays one move and notifies the state
	Listen(play func (gameid string, move tictactoe.Move) (win, draw bool, err error) )
}
