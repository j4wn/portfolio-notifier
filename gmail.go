package main

import (
	"encoding/base64"

	log "github.com/sirupsen/logrus"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	gmail "google.golang.org/api/gmail/v1"
)

var (
	user      = "me"
	message   gmail.Message
	mailScope = "https://www.googleapis.com/auth/gmail.send"
)

func sendMail(sender string, recipient string, portfolioValue string) *gmail.Message {
	// Use oauth2.NoContext if there isn't a good context to pass in.
	ctx := context.Background()

	mailLog := log.WithFields(log.Fields{
		"file": "gmail.go",
	})

	mailLog.Infof("GMail Client: %s", mailScope)
	client, err := google.DefaultClient(ctx, mailScope)
	if err != nil {
		mailLog.Fatalf("Could not build client: %v", err)
	}

	service, err := gmail.New(client)
	if err != nil {
		mailLog.Fatalf("Unable to retrieve gmail Client: %v", err)
	}

	// Compose the message
	messageStr := []byte("From: " + sender + "\r\n" +
		"To: " + recipient + "\r\n" +
		"Subject: Portfolio Value: " + portfolioValue + "\r\n" +
		"\r\nMessage body goes here! This is a message body.\n" +
		"Something else something.\n" +
		"MOAR, done.")

	// Place messageStr into message.Raw in base64 encoded format
	message.Raw = base64.URLEncoding.EncodeToString(messageStr)

	// Send the message
	mailLog.WithFields(log.Fields{
		"from":  sender,
		"to":    recipient,
		"value": portfolioValue,
	}).Info("Sending Mail")

	resp, err := service.Users.Messages.Send(user, &message).Do()
	if err != nil {
		mailLog.Fatalf("Could not insert message: %v, %#v", err, err)
	} else {
		mailLog.Info("Message sent!")
	}

	return resp
}
