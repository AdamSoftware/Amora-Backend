package helper


import (
	"net/http"
	"os"
	"time"
	"fmt"
	"context"
	openai "github.com/sashabaranov/go-openai"
	"Amora/internal/Logging"
)

// I need a struct to hold the message so that it can be spit out by the API to the user
type OpenAiMessage struct {
	Message string `json:"message"`
}

type OpenAiClient struct {
	Client *openai.Client
}

func GetOpenAiClient() *openai.Client {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		Logging.Logs.Warnf("OPENAI_API_KEY environment variable is not set")
		return nil
	}

	config := openai.DefaultConfig(apiKey)
	config.HTTPClient = &http.Client{
		Timeout: 30 * time.Second,
	}

	client := openai.NewClientWithConfig(config)
	if client == nil {
		Logging.Logs.Errorf("Failed to create OpenAI client")
		return nil
	}

	fmt.Println("OpenAI client created successfully")
	return client
}

func SendOpenAiMessage(client *openai.Client, partnerName string, partnerNickname string) (*OpenAiMessage, error) {
   if client == nil {
     return nil, fmt.Errorf("OpenAI client is not initialized")
   }

   nickname := partnerNickname
   if nickname == "" {
     nickname = "no nickname"
   }

   prompt := fmt.Sprintf(
     "Hello, I want you to play a character: you are this person's loved one, you are in a relationship with them. Do not break character under any circumstances. Here is this person's name: %s and this is their nickname (if any): %s",
      partnerName,
      nickname,
    )

   ctx := context.Background()

   resp, err := client.CreateChatCompletion(
      ctx,
      openai.ChatCompletionRequest{
          Model: openai.GPT3Dot5Turbo,
          Messages: []openai.ChatCompletionMessage{
             {
              Role:    openai.ChatMessageRoleUser,
              Content: prompt,
             },
           },
          MaxTokens: 1000,
        },
    )

  if err != nil {
    Logging.Logs.Errorf("Error sending message to OpenAI: %v", err)
    return nil, err
  }

  if len(resp.Choices) == 0 {
   return nil, fmt.Errorf("no response from OpenAI")
  }

  return &OpenAiMessage{Message: resp.Choices[0].Message.Content}, nil
}










