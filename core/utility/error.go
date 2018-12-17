/*
========================================================================================================================
Name        : core/utility/error.go
Author      : Andrea Graziani
Description : This file includes some utility function about error managing.
========================================================================================================================
*/
package utility

// Defaults
const (
	InvalidInput = "ERROR: Invalid input"
)

// This function is used to check a panic error.
func CheckPanicError(e error) {
	if e != nil {
		panic(e)
	}
}
