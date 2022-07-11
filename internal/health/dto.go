package health

type Status string

const (
	StatusHealthy   Status = "healthy"
	StatusUnhealthy Status = "unhealthy"
)

type Response struct {
	Status Status `json:"status"`
	Error  string `json:"error,omitempty"`
}
