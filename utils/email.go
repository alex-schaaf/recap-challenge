package utils

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

// SendEmail sends an email using AWS Simple Email Service.
func SendEmail(subject, message string) error {
	sender := os.Getenv("EMAIL_SENDER")
	recipient := os.Getenv("EMAIL_RECEIVER")
	charSet := "UTF-8"

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("eu-central-1")},
	)

	svc := ses.New(sess)

	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			CcAddresses: []*string{},
			ToAddresses: []*string{
				aws.String(recipient),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Text: &ses.Content{
					Charset: aws.String(charSet),
					Data:    aws.String(message),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String(charSet),
				Data:    aws.String(subject),
			},
		},
		Source: aws.String(sender),
	}

	_, err = svc.SendEmail(input)

	if err != nil {
		return err
	}

	fmt.Println("Email Sent to address: " + recipient)

	return nil
}
