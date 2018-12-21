/*
========================================================================================================================
Name        : core/remote/subscription.go
Author      : Andrea Graziani
Description : This file includes RPC-function used to perform "Subscription" task.
========================================================================================================================
*/
package remote

import (
	"Word-Count/core"
	"Word-Count/core/utility/global"
	"net/rpc"
)

type Subscription struct{}

// This structure represent the input of "Subscription" task.
type SubscriptionInput struct {
	WorkerAddress string
}

// This structure represent the output of "Subscription" task.
type SubscriptionOutput struct {
	Data string
}

// Following function represents the published RPC routine used to perform "Subscription" task.
func (*Subscription) Execute(pInput SubscriptionInput, pOutput *SubscriptionOutput) error {

	var mError error
	var mWorkerClientObject *rpc.Client

	core.MutexSubscription.Lock()

	mWorkerClientObject, mError = rpc.Dial(core.DefaultNetwork, pInput.WorkerAddress)
	global.CheckError(mError)

	core.AvailableWorkersList[pInput.WorkerAddress] = mWorkerClientObject
	core.AvailableWorkersNumber++
	core.MutexSubscription.Unlock()
	return nil
}
