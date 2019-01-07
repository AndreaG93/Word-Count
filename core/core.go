/*
========================================================================================================================
Name        : core/core.go
Author      : Andrea Graziani
Description : This file includes some global system crucial information.
========================================================================================================================
*/
package core

import (
	"net/rpc"
	"sync"
)

// Constants
const (
	DefaultFileName                  = "file"
	BufferSize                       = 1024
	DefaultNetwork                   = "tcp"
	DefaultServerRPCAddress          = "localhost:3001"
	DefaultServerFileReceiverAddress = "localhost:3000"
)

// System Global
var MutexSubscription = &sync.Mutex{}
var AvailableWorkersList = map[string]*rpc.Client{}
var AvailableWorkersNumber = 0
