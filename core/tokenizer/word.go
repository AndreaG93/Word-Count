/*
========================================================================================================================
Name        : word.go
Author      : Andrea Graziani
Version     : 1.0
Date		: 14/12/2018
Description : This file includes some utility function to tokenize a plain text input.
========================================================================================================================
*/
package tokenizer

import (
	"Word-Count/core"
	"Word-Count/core/utility"
	"bufio"
	"strings"
)

// This function is used to tokenize a specified plain text file.
// If successful, this function returns a "pWorkerCardinality"-dimensional array containing pointer to
// hash tables.
//
// According to MapReduce programming model this function represents a "Map Procedure"
// with some "Reduce" and "Shuffle" characteristics.
func WordFromString(pInput string) ([]map[string]uint32, error) {

	var mNestedHashTable map[string]uint32
	var mHastTableIndex uint32
	var mError error

	// 1 - Initialization scanner...
	// ====================================================================== //
	mWordScanner := bufio.NewScanner(strings.NewReader(pInput))
	mWordScanner.Split(bufio.ScanWords)

	// 2 - Initialization Output
	// ====================================================================== //

	// Create an empty slice with "pWorkerCardinality" length...
	mHashTable := make([]map[string]uint32, core.WorkerCardinality)

	// Crete "pWorkerCardinality" nested hash tables...
	for mHastTableIndex = 0; mHastTableIndex < core.WorkerCardinality; mHastTableIndex++ {
		mHashTable[mHastTableIndex] = make(map[string]uint32)
	}

	// 3 - Scanning and tokenize...
	// ====================================================================== //
	for mWordScanner.Scan() {

		// Get a word from file and change all his Unicode letters to their lower case...
		mCurrentWord := strings.ToLower(mWordScanner.Text())

		// Getting hash table index...
		if mHastTableIndex, mError = utility.GetHashIndexFromString(mCurrentWord, core.WorkerCardinality); mError != nil {
			return nil, mError
		}

		// Select correct "mNestedHashTable"...
		mNestedHashTable = mHashTable[mHastTableIndex]
		// Check word occurrences from selected "mNestedHashTable"...
		mWordOccurrences := mNestedHashTable[mCurrentWord]

		// Insert current word into selected "mNestedHashTable" update word occurrences if necessary...
		if mWordOccurrences == 0 {
			mNestedHashTable[mCurrentWord] = 1
		} else {
			mNestedHashTable[mCurrentWord] = mWordOccurrences + 1
		}
	}

	// Error reporting
	return mHashTable, mWordScanner.Err()
}
