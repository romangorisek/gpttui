package networkservice

type ChatgptService struct{}

func NewChatgptService() *ChatgptService {
	return &ChatgptService{}
}

func (chatgptService ChatgptService) GetAllSessions() ([]string, error) {
	return []string{"sess 1", "sess 2"}, nil
}
