package main

import (
	"fmt"

	"gopkg.in/gomail.v2"
)

func main() {
	host := ""
	port := 0
	user := ""
	password := ""

	msg := gomail.NewMessage()
	msg.SetHeader("From", "from@gmail.com")
	msg.SetHeader("To", "to@gmail.com")
	msg.SetHeader("Subject", "test message")
	msg.SetBody("text/html", "<h1>Escreva-se no canal</h1>")

	dialer := gomail.NewDialer(host, port, user, password)

	if err := dialer.DialAndSend(msg); err != nil {
		fmt.Println("Não DEU CERTO")
	} else {
		fmt.Println("DEU CERTO")
	}
}
