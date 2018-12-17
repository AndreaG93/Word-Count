/*
========================================================================================================================
Name        : core/task/map.go
Author      : Andrea Graziani
Description : This file includes RPC-function used to perform "Map" task.
========================================================================================================================
*/
package task

import (
	"Word-Count/core/tokenizer"
)

type MapTask struct{}

// This structure represent the input of "Map" task.
type MapTaskInput struct {
	Data string
}

// This structure represent the output of "Map" task.
type MapTaskOutput struct {
	Data []map[string]uint32
}

// Following function represents the published RPC routine used to perform "Map" task.
func (x *MapTask) Execute(pInput MapTaskInput, pOutput *MapTaskOutput) error {

	if mOutput, mError := tokenizer.WordFromString(pInput.Data); mError != nil {
		return mError
	} else {
		pOutput.Data = mOutput
		return nil
	}
}
