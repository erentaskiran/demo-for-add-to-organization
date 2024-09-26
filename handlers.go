package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
)

func (app *application) HandleAddUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("asd")
	err := app.SendEmail(Sender, Recipient, Subject, HtmlBody, TextBody)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
func (app *application) SendEmail(sender, recipient, subject, htmlBody, textBody string) error {

	client := ses.NewFromConfig(app.Cfg)

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
