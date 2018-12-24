/*
========================================================================================================================
Name        : main/worker_test.go
Author      : Andrea Graziani
Description : TEST UNIT (NOT DOCKER ENVIRONMENT)
========================================================================================================================
*/
package test

import (
	"Word-Count/core"
	"Word-Count/core/task/remote"
	"Word-Count/core/utility/global"
	"fmt"
	"net"
	"net/rpc"
	"sync"
	"testing"
)

func Test_Worker(t *testing.T) {

	mWaitGroup := sync.WaitGroup{}
	mArray := [4]string{"localhost:1000", "localhost:1001", "localhost:1003", "localhost:1004"}

	mWaitGroup.Add(len(mArray))

	// Publish "MapTask" and "Reduce" RPC...
	// ====================================================================== //
	global.CheckError(rpc.Register(&remote.Map{}))
	global.CheckError(rpc.Register(&remote.Reduce{}))

	// Start worker...
	// ====================================================================== //
	for _, mWorkerAddress := range mArray {

		go func(pWorkerRPCAddress string) {

			var mError error
			var mClient *rpc.Client
			var input remote.SubscriptionInput
			var output remote.SubscriptionOutput
			var mListenerRPC net.Listener

			// PHASE 1 - Allocation "RPC Listener"...
			// ====================================================================== //
			mListenerRPC, mError = net.Listen(core.DefaultNetwork, pWorkerRPCAddress)
			global.CheckError(mError)

			// PHASE 2 - Execute a "Subscription" to main server...
			// ====================================================================== //
			mClient, mError = rpc.Dial(core.DefaultNetwork, core.DefaultServerRPCAddress)
			global.CheckError(mError)

			input.WorkerAddress = pWorkerRPCAddress

			mError = mClient.Call("Subscription.Execute", &input, &output)
			global.CheckError(mError)

			// Waiting requests...
			fmt.Printf("Worker: Waiting on %s\n", pWorkerRPCAddress)
			for {
				rpc.Accept(mListenerRPC)
			}

		}(mWorkerAddress)
	}

	// Wait Worker...
	mWaitGroup.Wait()
}
