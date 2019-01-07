![](https://img.shields.io/badge/Programming_Language-Go-blue.svg)
![](https://img.shields.io/badge/Release-1.0-blue.svg)
![](https://img.shields.io/badge/Status-Tested-green.svg)

# Word-Count
This repository contains a simple Go application to perform a distributed word count of a plain text file. This project is designed for educational purposes. 

## How to run

### How to build and run server and worker with Docker

Execute following commands to setup docker:

1. `cd $(go env GOPATH)/src`
2. `git clone https://github.com/AndreaG93/Word-Count`
3. `cd $(go env GOPATH)/src/Word-Count`
4. `sudo sh DockerSetup.sh`

#### Note:

Rememeber to create `$GOPATH/src` directory if it doesn't exist.

`DockerSetup.sh` is a script used to build needed Docker images and containers necessary to run application.
Is important to specify that all containers work with **HOST network** driver **ONLY**: host networking driver only works on Linux hosts, and is not supported on Docker for Mac, Docker for Windows, or Docker EE for Windows Server (see [documentation](https://docs.docker.com/network/host/)).

### How to build and run client

Execute following commands to build and run client into your machine.

1. `cd $(go env GOPATH)/src/Word-Count/main` 
2. `go build client.go`
2. `./client [path to a text file]`
