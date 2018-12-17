package subscriptions

type WorkerSubscriptionTask struct{}

// This structure represent the input of "WorkerSubscription" task.
type SubscriptionInput struct {
	WorkerAddress string
}

// This structure represent the output of "Map" task.
type SubscriptionOutput struct {
	Data []map[string]uint32
}
