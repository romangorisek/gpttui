package llmservice

type GptService struct{}

func (gpts *GptService) GetAllSessions() ([]string, error) {
	return []string{"sess 1", "sess 2"}, nil
}

func (gpts *GptService) GetSession(id string) (any, error) {
	return "this should be some session details struct", nil
}

func (gpts *GptService) SendQuestion(question string, id string) (any, string, error) {
	return "this should be some session details struct", "session_id", nil
}
