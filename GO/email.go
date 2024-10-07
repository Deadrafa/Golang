package main

import (
	"fmt"
	"log"
	"net/smtp"
)

func send_to_email(Email string) {
	smtpHost := "smtp.mail.ru"
	smtpPort := "587"
	smtpUsername := "ivan_test_user@mail.ru"
	smtpPassword := "ySR4VHssrfPe2jkiArGq"

	from := "ivan_test_user@mail.ru"
	to := []string{Email}

	subject := "Subject: Приветствие\n"
	body := "Здравствуйте!\n\nСпасибо что воспользовались нашим сайтом.\nВаш пароль: 2005"

	msg := []byte(subject + "\n" + body)

	auth := smtp.PlainAuth("", smtpUsername, smtpPassword, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, msg)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Письмо успешно отправлено!")
}
