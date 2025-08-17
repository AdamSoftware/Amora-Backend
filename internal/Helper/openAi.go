package helper

import (
	"Amora/internal/Logging"
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	openai "github.com/sashabaranov/go-openai"
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

func SendOpenAiMessage(
	client *openai.Client, 
	partnerName string, 
	partnerNickname string, 
	isBirthday bool) (*OpenAiMessage, error) {
   if client == nil {
     return nil, fmt.Errorf("OpenAI client is not initialized")
   }

   nickname := partnerNickname
   if nickname == "" {
     nickname = "no nickname"
   }


	prompt := fmt.Sprintf(
		"You're this person's significant other. Today is either their birthday (IsBirthday: %t) or anniversary. You're texting them a short, flirty, and heartfelt message. Do NOT break character. Their name is %s and their nickname is %s (if it's empty, ignore it).\n\n" +
		"Instructions:\n" +
		"- Message must be max 4 sentences.\n" +
		"- Keep it casual, unique, flirty, and human — like a real text message.\n" +
		"- Use emojis naturally.\n" +
		"- If it's their birthday, wish them a happy birthday. If it's an anniversary, do not mention birthday — just celebrate the anniversary.\n" +
		"- Use the nickname naturally if it exists (don’t mention the word 'nickname').\n" +
		"- You may write a short poem if it fits, but keep it short.\n" +
		"- Do not use generic phrases or greetings — be creative.\n" +
		"- Do not mention AI, programming, or use any formatting like \\n.\n" +
		"- Avoid gender-specific language for example nothing say guy,girl,he him, she her.\n" +
		"- End with something like 'can't wait to see you soon' or something cute or flirty.\n" +
		"- you can be a little bit spicy but just don't go too explicit.\n" +
		"- remember you are human so you can use some slang words, and emojis",
		isBirthday,
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

	cleanMessage := strings.ReplaceAll(resp.Choices[0].Message.Content, "\n", " ")
	return &OpenAiMessage{Message: cleanMessage}, nil
}










