/*
========================================================================================================================
Name        : core/task/wordcount.go
Author      : Andrea Graziani
Description : This file function used to perform a "WordCount" task.
========================================================================================================================
*/
package remote

import (
	"Word-Count/core"
	"Word-Count/core/utility/file"
	"Word-Count/core/utility/global"
	"fmt"
	"io/ioutil"
	"net/rpc"
	"os"
	"path/filepath"
	"sync"
)

type WordCount struct{}

// This structure represent the input of "WordCount" task.
type WordCountInput struct {
	FileHash string
}

// This structure represent the output of "WordCount" task.
type WordCountOutput struct {
	Data map[string]uint32
}

// This function represents server "RequestTask-Task" task.
func (x *WordCount) Execute(pInput WordCountInput, pOutput *WordCountOutput) error {

	var mError error
	var mWorkingDirectory string
	var index int

	// Lock Worker Subscription and defer its unlock...
	// ====================================================================== //
	core.MutexSubscription.Lock()
	defer core.MutexSubscription.Unlock()

	// Allocation needed structures...
	// ====================================================================== //
	mWorkingDirectory = filepath.Join(os.TempDir(), pInput.FileHash)

	mMapInputArray := make([]MapInput, core.AvailableWorkersNumber)         // Input for "Map" task.
	mMapOutputArray := make([]MapOutput, core.AvailableWorkersNumber)       // Output generated by "Worker" after "Map" task.
	mReduceInputArray := make([]ReduceInput, core.AvailableWorkersNumber)   // Input for "Reduce" task.
	mReduceOutputArray := make([]ReduceOutput, core.AvailableWorkersNumber) // Output generated by "Worker" after "Reduce" task.

	for x := 0; x < core.AvailableWorkersNumber; x++ {
		mReduceInputArray[x].Data = make([]map[string]uint32, core.AvailableWorkersNumber)
		mReduceInputArray[x].Modulus = core.AvailableWorkersNumber
	}

	for x := 0; x < core.AvailableWorkersNumber; x++ {
		mReduceOutputArray[x].Data = make(map[string]uint32)
	}

	// Split received file...
	// ====================================================================== //
	mError = file.SplitByWord(mWorkingDirectory, core.DefaultFileName, core.AvailableWorkersNumber)
	global.CheckError(mError)

	// Perform "MAP" phase...
	// ====================================================================== //
	index = 0
	mWaitGroup := sync.WaitGroup{}              // A WaitGroup waits for a collection of goroutines to finish.
	mWaitGroup.Add(core.AvailableWorkersNumber) // Add delta to the "mWaitGroup" counter.

	for _, mWorker := range core.AvailableWorkersList {

		// Getting "FileName"
		mFileName := fmt.Sprintf("%s%s%s%d", mWorkingDirectory, string(os.PathSeparator), core.DefaultFileName, index)

		// Launch goroutine...
		go func(pWaitGroup *sync.WaitGroup, pWorker *rpc.Client, mFileName string, input *MapInput, output *MapOutput) {

			// Defer decrement of WaitGroup
			// counter by one...
			// ------------------------------- //
			defer pWaitGroup.Done()

			// Open input file and defer its
			// close...
			// ------------------------------- //
			mFile, mError := os.OpenFile(mFileName, os.O_RDONLY, 0666)
			global.CheckError(mError)
			defer func() {
				global.CheckError(mFile.Close())
			}()

			// Read input file...
			// ------------------------------- //
			mByteRead, mError := ioutil.ReadAll(mFile)
			global.CheckError(mError)

			// Populate "MapInput" structure...
			// ------------------------------- //
			input.Data = string(mByteRead)
			input.Modulus = core.AvailableWorkersNumber

			// Call RPC-Map...
			// ------------------------------- //
			global.CheckError(pWorker.Call("Map.Execute", &input, &output))

		}(&mWaitGroup, mWorker, mFileName, &mMapInputArray[index], &mMapOutputArray[index])

		// Update index...
		index++
	}

	// Waiting...
	mWaitGroup.Wait()

	// Perform "SHUFFLE" phase...
	// ====================================================================== //
	shuffle(mMapOutputArray, mReduceInputArray)

	// Perform "REDUCE" phase 1 (Performed by "Workers")...
	// ====================================================================== //
	index = 0
	mWaitGroup.Add(core.AvailableWorkersNumber) // Add delta to the "mWaitGroup" counter.

	for _, mWorker := range core.AvailableWorkersList {

		// Launch goroutine...
		go func(pWaitGroup *sync.WaitGroup, pWorker *rpc.Client, input *ReduceInput, output *ReduceOutput) {

			// Defer decrement of WaitGroup
			// counter by one...
			// ------------------------------- //
			defer pWaitGroup.Done()

			// Call RPC-Reduce...
			// ------------------------------- //
			global.CheckError(mWorker.Call("Reduce.Execute", &input, &output))

		}(&mWaitGroup, mWorker, &mReduceInputArray[index], &mReduceOutputArray[index])

		// Update index...
		index++
	}

	// Waiting...
	mWaitGroup.Wait()

	// Perform "REDUCE" phase 2 (Performed by Server)...
	// ====================================================================== //
	mFinalOutput := make(map[string]uint32)

	for i := 0; i < core.AvailableWorkersNumber; i++ {
		for key, value := range mReduceOutputArray[i].Data {
			mFinalOutput[key] = value
		}
	}

	// Populate output and exit...
	pOutput.Data = mFinalOutput
	return nil
}

// This function is used to perform "Server" initialization.
//
// According to MapReduce programming model this function represents a "Map Procedure"
// with some "Reduce" and "Shuffle" characteristics.
func shuffle(pInput []MapOutput, pOutput []ReduceInput) {

	for x := 0; x < core.AvailableWorkersNumber; x++ {
		for y := 0; y < core.AvailableWorkersNumber; y++ {
			pOutput[y].Data[x] = pInput[x].Data[y]
		}
	}
}
