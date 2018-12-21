/*
========================================================================================================================
Name        : core/task/map.go
Author      : Andrea Graziani
Description : This file includes RPC-function used to perform "Map" task.
========================================================================================================================
*/
package remote

import (
	"Word-Count/core/task/tokenization"
	"fmt"
)

type Map struct{}

// This structure represent the input of "Map" task.
type MapInput struct {
	Data    string
	Modulus int
}

// This structure represent the output of "Map" task.
type MapOutput struct {
	Data []map[string]uint32
}

// Following function represents the published RPC routine used to perform "Map" task.
func (x *Map) Execute(pInput MapInput, pOutput *MapOutput) error {

	if mOutput, mError := tokenization.WordFromString(pInput.Data, pInput.Modulus); mError != nil {
		return mError
	} else {
		pOutput.Data = mOutput
		fmt.Println(pOutput.Data)
		return nil
	}

}
