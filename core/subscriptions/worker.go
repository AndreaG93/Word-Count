package subscriptions

type Worker struct{}

// This structure represent the input of "WorkerSubscription" task.
type WorkerSubscriptionInput struct {
	WorkerAddress string
}

// This structure represent the output of "Map" task.
type WorkerSubscriptionOutput struct {
	Data []map[string]uint32
}

func (*Worker) Execute(pInput WorkerSubscriptionInput, pOutput *WorkerSubscriptionOutput) error {
	return nil
}
