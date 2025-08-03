package helper 

import (
	twilio "github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
	"os"
)

func SendResetSMS(phoneNumber, resetLink string) error {

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: os.Getenv("TWILIO_SID"),
		Password: os.Getenv("TWILIO_AUTH_TOKEN"),
	})

	params := &openapi.CreateMessageParams{}
	params.SetTo(phoneNumber)
	params.SetFrom(os.Getenv("TWILIO_PHONE_NUMBER"))
	params.SetBody("Reset your password here: " + resetLink)

	_, err := client.Api.CreateMessage(params)
	return err
}
