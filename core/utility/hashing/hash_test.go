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
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_WriteFileContainsOutput(t *testing.T) {

	var mOutput string
	var mError error
	var mAssert *assert.Assertions

	// Allocation "assert" object...
	mAssert = assert.New(t)

	// Getting Hash-Digest...
	mOutput, mError = GetHashFromFile("../../../_data.txt")
	fmt.Println(mOutput)
	mAssert.Nil(mError)
	mAssert.NotEmpty(mOutput)
}

/*
// Test for "getHashTableIndex" function.
func Test_getHashTableIndex(t *testing.T) {

	var mError error
	var mAssert *assert.Assertions

	// Allocation "assert" object...
	mAssert = assert.New(t)

	// Checking error...
	_, mError = GetHashIndexFromString("andrea", uint32(core.WorkerCardinality()))
	mAssert.Nil(mError)

	_, mError = GetHashIndexFromString("graziani", core.WorkerCardinality())
	mAssert.Nil(mError)

	_, mError = GetHashIndexFromString("valeria", core.WorkerCardinality())
	mAssert.Nil(mError)
}
*/
