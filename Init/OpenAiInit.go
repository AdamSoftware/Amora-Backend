package Init

import (
	"Amora/internal/Logging"
	"Amora/internal/helper"
	"fmt"
)

func InitOpenAi() *helper.OpenAiClient {
    client := helper.GetOpenAiClient()
    if client == nil {
        Logging.Logs.Warnf("OpenAI client is not initialized")
        return nil
    }
		fmt.Println("OpenAI client initialized successfully")
    return &helper.OpenAiClient{Client: client}
}



