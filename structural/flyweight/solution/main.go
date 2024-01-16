package main

import "fmt"

type ChatMessage struct {
	Content string
	sender  *Sender
}

type Sender struct {
	Name   string
	Avatar []byte
}

type SenderFactory struct {
	cacheSender map[string]*Sender
}

func (sf *SenderFactory) GetSender(name string) *Sender {
	return sf.cacheSender[name]
}

func main() {
	senderFactory := SenderFactory{cacheSender: map[string]*Sender{
		"ThinhPhan": {
			Name:   "ThinhPhan",
			Avatar: make([]byte, 1024*300), // 300kb,
		},
		"TPV": {
			Name:   "TPV",
			Avatar: make([]byte, 1024*400), // 400kb
		},
	}}

	fmt.Println([]ChatMessage{
		{
			Content: "hi",
			sender:  senderFactory.GetSender("ThinhPhan"),
		},
		{
			Content: "oh here you are",
			sender:  senderFactory.GetSender("TPV"),
		},
		{
			Content: "how are you doing?",
			sender:  senderFactory.GetSender("ThinhPhan"),
		},
		{
			Content: "I'm doing well?",
			sender:  senderFactory.GetSender("TPV"),
		},
	})
	// Total memory of avatars: 700kb
}
