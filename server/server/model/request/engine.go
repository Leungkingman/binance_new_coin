package request

type MainQueueTaskData struct {
	Ips []string `json:"ips"`
}

type QueueTaskStatus struct {
	QueueTaskRunning bool `json:"queueTaskRunning"`
}
