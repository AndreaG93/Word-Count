/*
========================================================================================================================
Name        : core/utility/hash_test.go
Author      : Andrea Graziani
Description : TEST SUITE
========================================================================================================================
*/
package hashing

import (
	"fmt"
	"os"
	"testing"
)

func Test_WriteFileContainsOutput(t *testing.T) {

	var mOutput string
	var mError error

	// Getting Hash-Digest...
	mOutput, mError = GetHashFromFile("_data_test.txt")
	if mError != nil {
		os.Exit(1)
	}

	fmt.Println(mOutput)
}
