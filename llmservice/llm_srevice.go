package llmservice

import (
	"errors"

	"github.com/sirupsen/logrus"
)

type LlmServiceInterface interface {
	GetAllSessions() ([]string, error)
	GetSession(id string) (any, error)                            // TODO: replace any with SessionHistory struct
	SendQuestion(question string, id string) (any, string, error) // TODO: replace any with SessionHistory struct, return session id as second param, useful when starting a new session
}

type LlmService struct {
	llmModel LlmServiceInterface
	apiKey   string
	log      *logrus.Logger
}

func New(log *logrus.Logger, model string, apiKey string) (*LlmService, error) {
	ls := &LlmService{apiKey: apiKey, log: log}
	switch model {
	case "GPT_4":
		ls.log.Info("gpt 4 model")
		ls.llmModel = &GptService{}
		return ls, nil
	case "CLAUD":
		ls.log.Info("CLAUD model")
		return nil, nil
	default:
		return nil, errors.New("LLM model not supported")
	}
}

func (ls *LlmService) Test() {
	ls.log.Infof("we have llmservice set, the model is: %T and the api key is %v\n", ls.llmModel, ls.apiKey)
	sessions, _ := ls.llmModel.GetAllSessions()
	ls.log.Infof("%+v\n", sessions)
}

func (ls *LlmService) GetAllSessions() ([]string, error) {
	return ls.llmModel.GetAllSessions()
}

func (ls *LlmService) GetSession(id string) (any, error) { // TODO: replace any with SessionHistory struct
	return ls.llmModel.GetSession(id)
}

func (ls *LlmService) SendQuestion(question string, id string) (any, string, error) { // TODO: replace any with SessionHistory struct, return session id as second param, useful when starting a new session
	return ls.llmModel.SendQuestion(question, id)
}
