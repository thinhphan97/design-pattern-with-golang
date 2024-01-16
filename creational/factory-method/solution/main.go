package main

import "fmt"

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

type NotificationService struct {
	notifier Notifier
}

func (s NotificationService) SendNotification(message string) {
	s.notifier.Send(message)
}

func CreateNotifier(t string) Notifier {
	if t == "sms" {
		return SMSNotifier{}
	}
	return EmailNotifier{}
}

func main() {
	s := NotificationService{
		notifier: CreateNotifier("sms"),
	}
	s.SendNotification("Hello world")
}
