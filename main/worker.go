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

	go task.WorkerInitialization(0)
	go task.WorkerInitialization(1)
	go task.WorkerInitialization(2)
	go task.WorkerInitialization(3)
	go task.WorkerInitialization(4)
	go task.WorkerInitialization(5)
	go task.WorkerInitialization(6)
	go task.WorkerInitialization(7)
	go task.WorkerInitialization(8)
	go task.WorkerInitialization(9)

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
