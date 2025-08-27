package messageservice

import (
	"Amora/internal/Logging"
	"Amora/internal/messages/messagerepository"
	"errors"
)

type MessageServiceInit struct {
	messageRepo *messagerepository.MessageRepositoryInit
}

func NewMessageServiceInit( messageRepo *messagerepository.MessageRepositoryInit) (*MessageServiceInit, error) {

	if messageRepo == nil {
		Logging.Logs.Errorf("user repository is nil")
		return nil, errors.New("user repository is nil")
	}

	return &MessageServiceInit{
		messageRepo: messageRepo,
	}, nil
}
