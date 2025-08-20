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
	isBirthday bool,
	SpicyContent int) (*OpenAiMessage, error) {
   if client == nil {
     return nil, fmt.Errorf("OpenAI client is not initialized")
   }

   nickname := partnerNickname
   if nickname == "" {
     nickname = "no nickname"
   }

	 // NOTE:spicy content is a number that will be from 0 - 2 
	 // 0 = clean, 1 = a little spicy, 2 = M rated content


	
	
	prompt := fmt.Sprintf(
		"You're texting your significant other for their birthday or anniversary, writing a flirty, heartfelt, and playful message that feels like a real human text. Never break character; you are their loving, craving partner. Use their name (%s) and their nickname (%s, only if provided). The content must strictly follow the chosen SpicyContent level (0, 1, or 2), using the corresponding style and explicitness.\n\n" +

		"Key instructions and flow:\n" +
		"- Chain-of-thought: First, consider the occasion (birthday or anniversary: %t) and how to reference it in a creative, playful way.\n" +
		"- Next, factor in the SpicyContent level and nickname rules.\n" +
		"- Brainstorm 2-3 message concepts, select the most unique and playful one.\n" +
		"- Only then generate your final text message.\n\n" +

		"Detailed requirements:\n" +
		"- Message: Maximum 4 sentences.\n" +
		"- Style: Real, unfiltered, human — use slang, emojis, and dirty talk as appropriate.\n" +
		"- NEVER mention AI, programming, or formatting (avoid \"\\n\").\n" +
		"- Include the nickname naturally if it exists, never refer to \"nickname\" directly.\n" +
		"- NEVER use gendered words (he/she/guy/girl/etc).\n" +
		"- Use emojis naturally and playfully.\n" +
		"- If birthday: must wish happy birthday. If anniversary: never mention birthday; celebrate the relationship instead.\n" +
		"- Uniqueness: Do NOT use generic greetings/phrases. Avoid clichés.\n" +
		"- Can include a short poem if it fits — still intimate and brief.\n" +
		"- End with something flirty, dirty, or teasing to build anticipation.\n" +
		"- Output must strictly match the SpicyContent level:\n" +
		"  - Level 0: Sweet, romantic, and clean. No innuendo or suggestive content.\n" +
		"  - Level 1: Light, playful innuendo, mildly spicy, but never explicit.\n" +
		"  - Level 2: Full-on erotic, bold, and explicit. Use vivid, dirty, unfiltered language.\n" +
		"- Never clean up or censor explicit content at level 2; always go as bold as allowed.\n" +
		"- Your persona is their loving, desirous partner — never break character, never include AI disclaimers.\n\n" +

		"Output Format:\n" +
		"- Output ONLY the final text message, nothing else. Do NOT include instructions, labels, commentary, or reasoning steps in your response — just the text, as if sending it directly.\n" +
		"- Maximum length: 4 sentences.\n" +
		"- No code blocks or markdown formatting. Just plain text suitable for SMS.\n\n" +

		"SpicyContent Level: %d\n",
		partnerName,
		nickname,
		isBirthday,
		SpicyContent,
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










