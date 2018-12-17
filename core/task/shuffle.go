/*
========================================================================================================================
Name        : core/task/wordcount.go
Author      : Andrea Graziani
Description : This file function used to perform a "Word-Count" task.
========================================================================================================================
*/
package task

import "Word-Count/core"

// This function is used to perform "Server" initialization.
//
// According to MapReduce programming model this function represents a "Map Procedure"
// with some "Reduce" and "Shuffle" characteristics.
func Shuffle(pInput []MapTaskOutput, pOutput []ReduceTaskInput) {

	for x := 0; x < core.WorkerCardinality; x++ {
		for y := 0; y < core.WorkerCardinality; y++ {
			pOutput[y].Data[x] = pInput[x].Data[y]
		}
	}
}
