/*
========================================================================================================================
Name        : main/map.go
Author      : Andrea Graziani
Description : This file includes "Worker Entry-Point"
========================================================================================================================
*/
package main

import (
	"Word-Count/core/task"
	"sync"
	"time"
)

// Worker "Entry-Point"
func main() {

	mWaitGroup := sync.WaitGroup{}
	mWaitGroup.Add(1)

	go task.WorkerInitialization("localhost:1000")
	time.Sleep(time.Second * 1)
	go task.WorkerInitialization("localhost:1001")
	time.Sleep(time.Second * 1)
	go task.WorkerInitialization("localhost:1002")
	time.Sleep(time.Second * 3)
	go task.WorkerInitialization("localhost:1003")
	time.Sleep(time.Second * 1)
	go task.WorkerInitialization("localhost:1004")

	mWaitGroup.Wait()
	/*
		var mError error
		var mWorkerAddress string

		mWorkerAddress, mError = os.Getenv("WORKER_ADDRESS")
		utility.CheckPanicError(mError)

		// Start...
		task.WorkerInitialization(mWorkerAddress)
	*/
}
