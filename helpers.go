package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
)

type Helpers struct {
	Cfg *aws.Config
}

func NewHelper(cfg *aws.Config) *Helpers {
	return &Helpers{Cfg: cfg}
}
func (helper *Helpers) SendEmail(recipient string) error {
	const (
		sender   = "erentskrn7@gmail.com"
		subject  = "asdasdasdasd"
		htmlBody = "<h1 style='color:blue'>Amazon SES Test Email (AWS SDK for Go)</h1><p>This email was sent with " +
			"<a href='https://aws.amazon.com/ses/'>Amazon SES</a> using the " +
			"<a href='https://aws.amazon.com/sdk-for-go/'>AWS SDK for Go</a>.</p>"
		textBody = "This email was sent with Amazon SES using the AWS SDK for Go."
		charSet  = "UTF-8"
	)

	client := ses.NewFromConfig(*helper.Cfg)

	input := &ses.SendEmailInput{
		Destination: &types.Destination{
			ToAddresses: []string{
				recipient,
			},
		},
		Message: &types.Message{
			Body: &types.Body{
				Html: &types.Content{
					Charset: aws.String("UTF-8"),
					Data:    aws.String(htmlBody),
				},
				Text: &types.Content{
					Charset: aws.String("UTF-8"),
					Data:    aws.String(textBody),
				},
			},
			Subject: &types.Content{
				Charset: aws.String("UTF-8"),
				Data:    aws.String(subject),
			},
		},
		Source: aws.String(sender),
	}

	_, err := client.SendEmail(context.TODO(), input)
	if err != nil {
		return fmt.Errorf("e-posta gönderilemedi: %v", err)
	}

	return nil
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func JSONError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(ErrorResponse{Error: message})
}

func JSONResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
