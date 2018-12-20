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
)

// Worker "Entry-Point"
func main() {

	mWaitGroup := sync.WaitGroup{}
	mWaitGroup.Add(1)

	go task.WorkerInitialization("localhost:1000")
	go task.WorkerInitialization("localhost:1001")
	go task.WorkerInitialization("localhost:1002")
	go task.WorkerInitialization("localhost:1003")

	mWaitGroup.Wait()
	/*
		var mError error
		var mWorkerID int

		mWorkerID, mError = strconv.Atoi(os.Getenv("WORKER_ID"))
		utility.CheckPanicError(mError)

		// Start...
		task.WorkerInitialization(mWorkerID)
	*/
}
