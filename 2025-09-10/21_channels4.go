package main

import (
	"fmt"
	"time"
)

func emailSender(emailChan chan string, done chan bool) {
	defer func() { done <- true }()

	for email := range emailChan {
		fmt.Println("Sending email to", email)
		time.Sleep(time.Second)
	}
}

func main() {
	emailChan := make(chan string, 10)
	done := make(chan bool)

	go emailSender(emailChan, done)

	for i := 0; i < 5; i++ {
		emailChan <- fmt.Sprintf("%d@example.com", i)
	}

	close(emailChan)
	<-done
}
