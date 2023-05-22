package data

// PriorityQueueItem defines the items of the priority queue
type PriorityQueueItem struct {
	Priority float64
	ItemName string
	Message  string
	Url      string
}

// ScheduleItem defines the request body for /schedule
type ScheduleItem struct {
	Message string  `json:"message"`
	Url     string  `json:"url"`
	Period  float64 `json:"period"`
	Name    string  `json:"name"`
}

// ScheduleResponse defines the response body for /schedule
type ScheduleResponse struct {
	StatusCode int `json:"statusCode"`
}
