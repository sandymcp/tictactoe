# tictactoe
Simple console implemementation and web interfaces

# Web
To play the game we need to share the state between the two players
and transfer it to the players apps/browsers, where the rendering will
be done.

We need to identify each game uniquely and store it in a datastore,
with something as simple as tictactoe we don't need true persistence we
can use something like REDIS, with a more complex time consuming game like
chess we probably want to use a real database Cassandra, Mongo etc. We
don't want to bind any state into our server program, so that we can scale up.
by simply adding more server processes.

Each game needs a unique identifier so that it can be found on the server and
the moves can be sent back to the servers

To notify the players with new state when a move is played, a push mechanism to notify
 the player is needed e.g. WebSockets.
 
We need functionality for a player to register and invite friends, who are registered in the
system. 

##Cleanup
When a game is finished (win or draw) the (transient) game needs to be cleaned up, equally
the same thing should happen if one of the connections to a client is lost.