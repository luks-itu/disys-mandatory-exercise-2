# disys-mandatory-exercise-2
Distributed Mutual Exclusion

## Starting the program
First, start the server. Then start any number of clients you want.
They should be run from different terminal instances.

### Server
Run `go run .` in the `/server` folder.
Then input your desired host address.

### Client
Run `go run .` in the `/client` folder.
Then input the address of the server.

## Stopping the program
Both the server and the clients are stopped by pressing _Enter_ in their terminals.
Logging does not always work correctly if you use _ctrl + c_.
