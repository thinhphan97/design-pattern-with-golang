package main

import "fmt"

type ChatMessage struct {
	Content      string
	SenderName   string
	SenderAvarta []byte
}

func main() {
	fmt.Println([]ChatMessage{
		{
			Content:      "hi",
			SenderName:   "ThinhPhan",
			SenderAvarta: make([]byte, 1024*300), // 300kb
		},
		{
			Content:      "oh here you are",
			SenderName:   "TPV",
			SenderAvarta: make([]byte, 1024*400), // 400kb
		},
		{
			Content:      "how are you doing?",
			SenderName:   "ThinhPhan",
			SenderAvarta: make([]byte, 1024*300), // 300kb
		},
		{
			Content:      "I'm doing well?",
			SenderName:   "TPV",
			SenderAvarta: make([]byte, 1024*400), // 400kb
		},
	})
	// We just have Peter and Mary as the senders. Everytime they send a new message, a lot of memory being wasted
	// to represent their avatars.

	// Total memory of avatars: 300 + 400 + 300 + 400 = 1,4kb
}
