package main

import "fmt"

type NotificationService struct {
	notifierType string
}

func (s NotificationService) SendNotification(message string) {
	if s.notifierType == "email" {
		fmt.Printf("Sending messange: %s (sender email)\n", message)
	} else if s.notifierType == "sms" {
		fmt.Printf("Sending message: %s (sender sms)\n", message)
	}

}

func main() {
	s := NotificationService{notifierType: "email"}
	s.SendNotification("Hello world")
}
