/*
========================================================================================================================
Name        : core/task/reduce.go
Author      : Andrea Graziani
Description : This file includes RPC-function used to perform "Reduce" task.
========================================================================================================================
*/
package remote

type Reduce struct{}

// This structure represent the input of "Reduce" task.
type ReduceInput struct {
	Data    []map[string]uint32
	Modulus int
}

// This structure represent the output of "Reduce" task.
type ReduceOutput struct {
	Data map[string]uint32
}

// Following function represents the published RPC routine used to perform "Reduce" task.
func (x *Reduce) Execute(pInput ReduceInput, pOutput *ReduceOutput) error {

	pOutput.Data = make(map[string]uint32)

	for j := 0; j < pInput.Modulus; j++ {
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
