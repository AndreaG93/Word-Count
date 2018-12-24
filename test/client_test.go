/*
========================================================================================================================
Name        : main/client.go
Author      : Andrea Graziani
Description : TEST UNIT (NOT DOCKER ENVIRONMENT)
========================================================================================================================
*/
package test

import (
	"Word-Count/core/task"
	"testing"
)

func Test_Client(t *testing.T) {
	task.ClientInitialization("_data_test.txt")
}
