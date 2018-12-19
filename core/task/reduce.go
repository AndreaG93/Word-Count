/*
========================================================================================================================
Name        : core/task/reduce.go
Author      : Andrea Graziani
Description : This file includes RPC-function used to perform "Reduce" task.
========================================================================================================================
*/
package task

import (
	"Word-Count/core"
)

type Reduce struct{}

// This structure represent the input of "Reduce" task.
type ReduceTaskInput struct {
	Data []map[string]uint32
}

// This structure represent the output of "Reduce" task.
type ReduceTaskOutput struct {
	Data map[string]uint32
}

// Following function represents the published RPC routine used to perform "Reduce" task.
func (x *Reduce) Execute(pInput ReduceTaskInput, pOutput *ReduceTaskOutput) error {

	pOutput.Data = make(map[string]uint32)

	for j := 0; j < core.WorkerCardinality; j++ {
		for key, value := range pInput.Data[j] {

			// Check word occurrences from selected "mNestedHashTable"...
			mWordOccurrences := pOutput.Data[key]

			// Insert current word into selected "mNestedHashTable" update word occurrences if necessary...
			if mWordOccurrences == 0 {
				pOutput.Data[key] = value
			} else {
				pOutput.Data[key] = mWordOccurrences + value
			}
		}

	}
	return nil
}
