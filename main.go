package main

import (
	"fmt"

	"gopkg.in/gomail.v2"
)

func main() {
	host := "sandbox.smtp.mailtrap.io"
	port := 587
	user := "fb52525f87cad8"
	password := "3c8c996b2c1c9b"

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
