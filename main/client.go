/*
========================================================================================================================
Name        : main/client.go
Author      : Andrea Graziani
Description : This file includes "Client Entry-Point"
========================================================================================================================
*/
package main

import (
	"Word-Count/core/task"
	"fmt"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Printf("USAGE: %s [FILEPATH]", os.Args[0])
		os.Exit(1)
	}

	task.ClientInitialization(os.Args[1])
}
