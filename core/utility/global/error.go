/*
========================================================================================================================
Name        : core/utility/error.go
Author      : Andrea Graziani
Description : This file includes some utility function about error managing.
========================================================================================================================
*/
package global

// Defaults
const (
	InvalidInput = "ERROR: Invalid input"
)

// This function is used to check a panic error.
func CheckError(e error) {
	if e != nil {
		panic(e)
	}
}
