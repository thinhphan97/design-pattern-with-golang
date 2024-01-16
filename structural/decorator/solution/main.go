package main

import (
	"fmt"
)

type Notifier interface {
	Send(message string)
}

type EmailNotifier struct{}

func (EmailNotifier) Send(message string) {
	fmt.Printf("Sending message: %s (sender: Email)\n", message)
}

type SMSNotifier struct{}

func (SMSNotifier) Send(message string) {
	fmt.Printf("Sending message: %s (sender: SMS)\n", message)
}

type TelegramNotifier struct{}

func (telegramNotifier TelegramNotifier) Send(message string) {
	fmt.Printf("Sending message: %s (sender: Telegram)\n", message)
}

type NotifierDecorator struct {
	core     *NotifierDecorator
	notifier Notifier
}

func (nd NotifierDecorator) Send(message string) {
	nd.notifier.Send(message)

	if nd.core != nil {
		nd.core.Send(message)
	}
}

func (nd NotifierDecorator) Decorator(notifier Notifier) NotifierDecorator {
	return NotifierDecorator{
		core:     &nd,
		notifier: notifier,
	}
}

func NewNotifierDecorator(notifier Notifier) NotifierDecorator {
	return NotifierDecorator{notifier: notifier}
}

type NotificationService struct {
	notifier Notifier
}

func (s NotificationService) SendNotification(message string) {
	s.notifier.Send(message)
}

func main() {
	notifier := NewNotifierDecorator(EmailNotifier{}).Decorator(SMSNotifier{}).Decorator(TelegramNotifier{})

	s := NotificationService{
		notifier: notifier,
	}
	s.SendNotification("Hello world")
}
