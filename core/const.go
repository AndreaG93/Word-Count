/*
========================================================================================================================
Name        : core/const.go
Author      : Andrea Graziani
Description : This file includes
========================================================================================================================
*/
package core

// Constants
const (
	DefaultFileName                  = "file"
	WorkerCardinality                = 9
	BufferSize                       = 1024
	DefaultRPCPort                   = "1234"
	DefaultNetwork                   = "tcp"
	DefaultServerFileReceiverAddress = "localhost:2000"
)

// Following dictionary contains a map between worker ID and his network address
var addressWorkerDictionary = map[int]string{
	0: "localhost:1000",
	1: "localhost:1001",
	2: "localhost:1002",
	3: "localhost:1003",
	4: "localhost:1004",
	5: "localhost:1005",
	6: "localhost:1006",
	7: "localhost:1007",
	8: "localhost:1008",
	9: "localhost:1009",
}

// This function is used to retrieve Worker Network Address from a specified Worker ID.
func GetWorkerAddress(pWorkerID int) string {
	return addressWorkerDictionary[pWorkerID]
}

// This function is used to read configuration file and retrieve server address.
func GetServerAddress() string {
	return "localhost:" + DefaultRPCPort
}
