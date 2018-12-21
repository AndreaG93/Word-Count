/*
========================================================================================================================
Name        : core/utility/hash.go
Author      : Andrea Graziani
Description : This file includes some utility function about hash algorithms.
========================================================================================================================
*/
package hashing

import (
	"Word-Count/core/utility/global"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"hash/fnv"
	"os"
)

// FNV-1a is a not cryptographic hash function:
// 1) Fast to compute and designed for fast hash table.
// 2) Slightly better avalanche characteristics than FNV-1 hash function.
var lHashAlgorithm = fnv.New32a()

// Secure Hash Algorithm...
var lCryptoHashAlgorithm = sha1.New()

// This function is used to compute a m-dimensional hash-table's index from an hash digest computed
// from a specified string.
func GetHashIndexFromString(pInput string, pModulus uint32) (uint32, error) {

	// Checking pInput
	if pInput == "" {
		return 0, errors.New(global.InvalidInput)
	}

	// Getting digest and defer reset of hash state...
	if _, mError := lHashAlgorithm.Write([]byte(pInput)); mError != nil {
		return 0, mError
	}
	defer lHashAlgorithm.Reset()
	return lHashAlgorithm.Sum32() % pModulus, nil
}

// This function is used to compute a crypto hash digest from first 256 byte of a specified file name.
func GetHashFromFile(pFilePath string) (string, error) {

	var mBuffer []byte
	var mError error
	var mInputFile *os.File

	// Checking input
	if pFilePath == "" {
		return "", errors.New(global.InvalidInput)
	}

	// Open specified input-file and defer its closing...
	if mInputFile, mError = os.OpenFile(pFilePath, os.O_RDONLY, 0666); mError != nil {
		return "", mError
	}
	defer func() {
		global.CheckError(mInputFile.Close())
	}()

	// Read first 256 byte from specified file...
	mBuffer = make([]byte, 256)
	if _, mError = mInputFile.Read(mBuffer); mError != nil {
		return "", mError
	}

	// Getting digest and defer reset of hash state...
	if _, mError := lCryptoHashAlgorithm.Write(mBuffer); mError != nil {
		return "", mError
	}
	defer lCryptoHashAlgorithm.Reset()
	return base64.URLEncoding.EncodeToString(lCryptoHashAlgorithm.Sum(nil)), nil
}
