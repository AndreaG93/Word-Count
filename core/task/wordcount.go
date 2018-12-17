/*
========================================================================================================================
Name        : core/task/wordcount.go
Author      : Andrea Graziani
Description : This file function used to perform a "WordCount" task.
========================================================================================================================
*/
package task

import (
	"Word-Count/core"
	"Word-Count/core/file"
	"Word-Count/core/utility"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type WordCountTask struct{}

// This structure represent the input of "WordCount" task.
type WordCountTaskInput struct {
	FileHash string
}

// This structure represent the output of "WordCount" task.
type WordCountTaskOutput struct {
	Data map[string]uint32
}

// This function represents server "RequestTask-Task" task.
func (x *WordCountTask) Execute(pInput WordCountTaskInput, pOutput *WordCountTaskOutput) error {

	var mError error
	var mWorkingDirectory string

	// Allocation needed structures...
	// ====================================================================== //
	mWorkingDirectory = filepath.Join(os.TempDir(), pInput.FileHash)

	mMapTaskInputArray := make([]MapTaskInput, core.WorkerCardinality)         // Input for "Map" task.
	mMapTaskOutputArray := make([]MapTaskOutput, core.WorkerCardinality)       // Output generated by "Worker" after "Map" task.
	mReduceTaskInputArray := make([]ReduceTaskInput, core.WorkerCardinality)   // Input for "Reduce" task.
	mReduceTaskOutputArray := make([]ReduceTaskOutput, core.WorkerCardinality) // Output generated by "Worker" after "Reduce" task.

	for x := 0; x < core.WorkerCardinality; x++ {
		mReduceTaskInputArray[x].Data = make([]map[string]uint32, core.WorkerCardinality)
	}

	for x := 0; x < core.WorkerCardinality; x++ {
		mReduceTaskOutputArray[x].Data = make(map[string]uint32)
	}

	// Split received file...
	// ====================================================================== //
	mError = file.SplitByWord(mWorkingDirectory, core.DefaultFileName, core.WorkerCardinality)
	utility.CheckPanicError(mError)

	// Perform "MAP" phase...
	// ====================================================================== //
	for i := 0; i < core.WorkerCardinality; i++ {

		// Opening split-file...
		// ------------------------------- //
		mFile, mError := os.OpenFile(fmt.Sprintf("%s%s%s%d", mWorkingDirectory, string(os.PathSeparator), core.DefaultFileName, i), os.O_RDONLY, 0666)
		utility.CheckPanicError(mError)

		// Read and close split-file...
		// ------------------------------- //
		mByteRead, mError := ioutil.ReadAll(mFile)
		utility.CheckPanicError(mError)
		utility.CheckPanicError(mFile.Close())

		// Populate input...
		// ------------------------------- //
		mMapTaskInputArray[i].Data = string(mByteRead)

		// Call RPC-Map method...
		// ------------------------------- //
		mError = workerRPCInterfaces[i].Call("MapTask.Execute", &(mMapTaskInputArray[i]), &mMapTaskOutputArray[i])
		utility.CheckPanicError(mError)
	}

	// Perform "SHUFFLE" phase...
	// ====================================================================== //
	Shuffle(mMapTaskOutputArray, mReduceTaskInputArray)

	// Perform "REDUCE" phase 1 (Performed by Worker)...
	// ====================================================================== //
	for i := 0; i < core.WorkerCardinality; i++ {

		mError = workerRPCInterfaces[i].Call("Reduce.Execute", &mReduceTaskInputArray[i], &mReduceTaskOutputArray[i])
		utility.CheckPanicError(mError)
	}

	// Perform "REDUCE" phase 2 (Performed by Server)...
	// ====================================================================== //
	mFinalOutput := make(map[string]uint32)

	for i := 0; i < core.WorkerCardinality; i++ {
		for key, value := range mReduceTaskOutputArray[i].Data {
			mFinalOutput[key] = value
		}
	}

	// Populate output and exit...
	pOutput.Data = mFinalOutput
	return nil
}
