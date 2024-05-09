package model

type HealthCheckResponse struct {
	Status string `json:"status"`
}

type MessageResponse struct {
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}
