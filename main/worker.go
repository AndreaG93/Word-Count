/*
========================================================================================================================
Name        : main/worker.go
Author      : Andrea Graziani
Description : This file includes "Worker Entry-Point"
========================================================================================================================
*/
package main

import (
	"Word-Count/core/task"
	"os"
)

func main() {
	task.WorkerInitialization(os.Getenv("WORKER_ADDRESS"))
}
