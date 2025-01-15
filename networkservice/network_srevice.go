package networkservice

type NetworkService interface {
	GetAllSessions() ([]string, error)
	// GetSessionHistory(id string) (SessionHistory, error)
}
