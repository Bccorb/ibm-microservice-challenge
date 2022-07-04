# The Pokemon Global Trade System Prototype
This application was created as a coding challenge to create a bi-directional gRPC application in Go.
The current state of this project represents a bi-directional client and server with the ability for clients
to submit Trade Requests to the server and confirm that there is a trade avalible to be made or not. 

## About this application
The Makefile can be used to run commands related to the building and running of this application. Avalible commands are desribed below: 

`make gen` - Generate the protobuf stubs for the client and server. 
`make clean` - remove the protobuf stubs from the code base. 
`make server` - Run the server side of the global trade system.
`make client` - Run the client side of the global trade system.
`make install` - Install all the required files.
`make test` - Run the unit test suite for the application.

## Docker
The application is equiped with it's own dockerfiles for running the client and server inside of a docker container. 

### Building docker image
To build a server image run: `docker build -f server.Dockerfile . -t gts-server`
To build a client image run: `docker build -f client.Dockerfile . -t gts-client`

## Testing
Testing is not complete code coverage, but does excerise the expectations from the server, Trade Requests, and Trade Responses. 

To run the test suite run: `go test ./tests`

## Improvements to this system
This is currently a prototype and given the time the next things I would seek to achieve are:

- Add a database for storing trade requests
- Docker compose to allow the two containers to talk to each other locally
- Removing trade requests from the database when a trade is found
- Integrate the level system to request and trade pokemon of desired level ranges
- More logging and persist the client
- Complete testing 


