# Word-Count
This repository contains a simple Go application to perform a distributed word count of a plain text file. This project is designed for educational purposes. 

## How to run

Execute following commands to setup docker:

1. `cd $(go env GOPATH)/src` 
2. `git clone https://github.com/AndreaG93/Word-Count`
3. `cd $(go env GOPATH)/src/Word-Count`
4. `sudo sh DockerSetup.sh`

(**Create $GOPATH/src directory if it doesn't exist.**)

Execute following commands to build and run client into your machine.

1. `cd $(go env GOPATH)/src/Word-Count/main` 
2. `go build client.go`
2. `./client [path to a text file]`
