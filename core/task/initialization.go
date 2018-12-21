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
	"Word-Count/core/task/remote"
	"Word-Count/core/utility/file"
	"Word-Count/core/utility/global"
	"Word-Count/core/utility/hashing"
	"fmt"
	"net"
	"net/rpc"
)

// This function is used to perform "Client" initialization.
func ClientInitialization(pFilePath string) {

	var mError error
	var mWordCountTaskInput remote.WordCountInput
	var mWordCountTaskOutput remote.WordCountOutput
	var mFileHash string
	var mClient *rpc.Client

	// PHASE 1 - Computing file-hash...
	// ====================================================================== //
	mFileHash, mError = hashing.GetHashFromFile(pFilePath)
	global.CheckError(mError)

	// PHASE 2 - Send file to server...
	// ====================================================================== //
	global.CheckError(file.Send(pFilePath, mFileHash, core.DefaultServerFileReceiverAddress))

	// PHASE 3 - Request "Word-Count" service through RPC
	// ====================================================================== //
	mClient, mError = rpc.Dial(core.DefaultNetwork, core.DefaultServerRPCAddress)
	global.CheckError(mError)

	// Is used by server to find sent file to compute...
	mWordCountTaskInput.FileHash = mFileHash

	mError = mClient.Call("WordCount.Execute", &mWordCountTaskInput, &mWordCountTaskOutput)
	global.CheckError(mError)

	// Print output
	fmt.Println(mWordCountTaskOutput.Data)
}

// This function is used to perform "Server" initialization.
func ServerInitialization() {

	var mServerListenerFileTransfer net.Listener
	var mServerListenerRPC net.Listener
	var mError error

	// PHASE 1 - Allocate needed "Listener" objects used for "File-Transfer"
	// 		     and RPC postponing their closing) ...
	// ====================================================================== //
	mServerListenerFileTransfer, mError = net.Listen(core.DefaultNetwork, core.DefaultServerFileReceiverAddress)
	global.CheckError(mError)
	defer func() {
		global.CheckError(mServerListenerFileTransfer.Close())
	}()

	mServerListenerRPC, mError = net.Listen(core.DefaultNetwork, core.DefaultServerRPCAddress)
	global.CheckError(mError)
	defer func() {
		global.CheckError(mServerListenerRPC.Close())
	}()

	// Publish "WorkCount" and "WorkerSubscribe" task...
	global.CheckError(rpc.Register(&remote.WordCount{}))
	global.CheckError(rpc.Register(&remote.Subscription{}))

	// PHASE 2 - Start listening for "File-Transfer"...
	// ====================================================================== //
	go func(x net.Listener) {
		for {
			fmt.Println("Server: Waiting File!")
			_, mError = file.Receive(x)
			global.CheckError(mError)
		}
	}(mServerListenerFileTransfer)

	// PHASE 3 - Start listening for RPC...
	// ====================================================================== //
	for {
		fmt.Println("Server: Waiting RPC!")
		rpc.Accept(mServerListenerRPC)
	}
}

// "WorkerSubscription" initialization task.
func WorkerInitialization(pWorkerRPCAddress string) {

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

	// PHASE 3 - Publish "MapTask" and "Reduce" RPC, create a TCP listener
	//			 and wait RPC request from server...
	// ====================================================================== //
	global.CheckError(rpc.Register(&remote.Map{}))
	global.CheckError(rpc.Register(&remote.Reduce{}))

	// Waiting requests...
	fmt.Printf("Worker: Waiting on %s\n", pWorkerRPCAddress)
	for {
		rpc.Accept(mListenerRPC)
	}
}
