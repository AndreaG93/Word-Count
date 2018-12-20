package subscriptions

import "Word-Count/core"

type Worker struct{}

// This structure represent the input of "WorkerSubscription" task.
type WorkerSubscriptionInput struct {
	WorkerAddress string
}

// This structure represent the output of "Map" task.
type WorkerSubscriptionOutput struct {
	Data string
}

func (*Worker) Execute(pInput WorkerSubscriptionInput, pOutput *WorkerSubscriptionOutput) error {

	core.SubscribeSystemWorker(pInput.WorkerAddress)
	return nil
}
