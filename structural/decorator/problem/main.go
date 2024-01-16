package main

import (
	"fmt"
)

type Notifier interface {
	Send(message string)
}

type EmailNotifier struct{}

func (EmailNotifier) Send(message string) {
	fmt.Printf("Sending message: %s (sender: Email)", message)
}

type SMSNotifier struct{}

func (SMSNotifier) Send(message string) {
	fmt.Printf("Sending message: %s (sender: SMS)", message)
}

// If I have more notifiers, I have to define more types to combine them!!!
type EmailSMSNotifier struct {
	amailNotifier EmailNotifier
	smsNotifier   SMSNotifier
}

func (notifier EmailSMSNotifier) Send(message string) {
	notifier.amailNotifier.Send(message)
	notifier.smsNotifier.Send(message)
}

type NotificationService struct {
	notifier Notifier
}

func (s NotificationService) SendNotification(message string) {
	s.notifier.Send(message)
}

func main() {
	s := NotificationService{
		notifier: EmailSMSNotifier{},
	}
	s.SendNotification("Hello world")
}
