/*
========================================================================================================================
Name        : core/task/initialization.go
Author      : Andrea Graziani
Description : This file includes initialization tasks.
========================================================================================================================
*/
package task

import (
	"Word-Count/core"
	"Word-Count/core/file"
	"Word-Count/core/subscriptions"
	"Word-Count/core/utility"
	"fmt"
	"net"
	"net/rpc"
)

var workerRPCInterfaces []*rpc.Client = nil // This array contains "*Client" object used to perform RPC calls.

// This function is used to perform allocation and initialization of "Worker-RPC-interfaces" that is
// an array of "*Client" object used to perform RPC calls.
func workerRPCInterfacesInitialization() {

	var mError error

	workerRPCInterfaces = make([]*rpc.Client, core.WorkerCardinality)

	for mIndex := range workerRPCInterfaces {
		workerRPCInterfaces[mIndex], mError = rpc.Dial(core.DefaultNetwork, core.GetWorkerAddress(mIndex))
		utility.CheckPanicError(mError)
	}
}

// This function is used to perform "Client" initialization.
func ClientInitialization(pFilePath string) {

	var mError error
	var mWordCountTaskInput WordCountTaskInput
	var mWordCountTaskOutput WordCountTaskOutput
	var mFileHash string
	var mClient *rpc.Client

	// PHASE 0 - Computing file-hash...
	// ====================================================================== //
	mFileHash, mError = utility.GetHashFromFile(pFilePath)
	utility.CheckPanicError(mError)

	// PHASE 1 - Send file to server...
	// ====================================================================== //
	mError = file.Send(pFilePath, mFileHash, core.DefaultServerFileReceiverAddress)
	utility.CheckPanicError(mError)

	// PHASE 2 - Preparing RPC call...
	// ====================================================================== //
	mClient, mError = rpc.Dial(core.DefaultNetwork, core.GetServerAddress())
	utility.CheckPanicError(mError)

	// Is used by server to find sent file to compute...
	mWordCountTaskInput.FileHash = mFileHash

	// PHASE 3 - Send RPC request...
	// ====================================================================== //
	mError = mClient.Call("WordCountTask.Execute", &mWordCountTaskInput, &mWordCountTaskOutput)
	utility.CheckPanicError(mError)

	// Print output
	fmt.Println(mWordCountTaskOutput.Data)
}

// This function is used to perform "Server" initialization.
func ServerInitialization() {

	var mServerListenerFileTransfer net.Listener
	var mServerListenerRPC net.Listener
	var mError error

	// PHASE 1 - Allocate needed "Listener" objects...
	// ====================================================================== //
	mServerListenerFileTransfer, mError = net.Listen(core.DefaultNetwork, core.DefaultServerFileReceiverAddress)
	utility.CheckPanicError(mError)
	defer func() {
		utility.CheckPanicError(mServerListenerFileTransfer.Close())
	}()

	mServerListenerRPC, mError = net.Listen(core.DefaultNetwork, core.GetServerAddress())
	utility.CheckPanicError(mError)
	defer func() {
		utility.CheckPanicError(mServerListenerRPC.Close())
	}()

	// Publish "WorkCount" and "WorkerSubscribe" task...
	mError = rpc.Register(&WordCountTask{})
	utility.CheckPanicError(mError)

	mError = rpc.Register(&subscriptions.Worker{})
	utility.CheckPanicError(mError)

	// PHASE 2 - Initialization "Worker-RPC-interfaces" to perform RPC call to worker...
	// ====================================================================== //
	//workerRPCInterfacesInitialization()

	// PHASE 3 - Start listening for file uploading...
	// ====================================================================== //
	go func(x net.Listener) {
		for {
			fmt.Println("Waiting File!")
			_, mError = file.Receive(x)
			utility.CheckPanicError(mError)
		}
	}(mServerListenerFileTransfer)

	// PHASE 4 - Start listening for "WorkCount" task request...
	// ====================================================================== //
	for {
		fmt.Println("Server Ready!")
		rpc.Accept(mServerListenerRPC)
	}
}

// "Worker" initialization task.
func WorkerInitialization(pWorkerID string) {

	var mError error
	var mClient *rpc.Client

	// PHASE 2 - Preparing RPC call...
	// ====================================================================== //
	mClient, mError = rpc.Dial(core.DefaultNetwork, core.GetServerAddress())
	utility.CheckPanicError(mError)

	var input subscriptions.WorkerSubscriptionInput
	var output subscriptions.WorkerSubscriptionOutput

	input.WorkerAddress = pWorkerID

	// PHASE 3 - Send RPC request...
	// ====================================================================== //
	mError = mClient.Call("Worker.Execute", &input, &output)
	utility.CheckPanicError(mError)

	/*
		var mError error
		var mListener net.Listener

		// Publish "MapTask" and "Reduce" service...
		mError = rpc.Register(&MapTask{})
		mError = rpc.Register(&Reduce{})
		//utility.CheckPanicError(mError)

		// Create a TCP listener that will listen on specified port...
		mListener, mError = net.Listen(core.DefaultNetwork, core.GetWorkerAddress(pWorkerID))
		utility.CheckPanicError(mError)

		// Waiting requests...
		for {
			fmt.Printf("Waiting on %s\n", core.GetWorkerAddress(pWorkerID))
			rpc.Accept(mListener)
		}
	*/
}
