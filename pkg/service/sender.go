package service

import (
	"Run_Hse_Run/pkg/mailer"
	"bytes"
	"html/template"
	"log"
	"math/rand"
)

type SenderService struct {
	sender *mailer.EmailSender
}

func (s *SenderService) SendEmail(email string) error {
	code := rand.Intn(9000) + 1000

	mu.Lock()
	Codes[email] = code
	mu.Unlock()

	buffer := bytes.NewBufferString("")
	tmpl, _ := template.ParseFiles("templates/message.html")
	err := tmpl.Execute(buffer, struct {
		Code int
	}{
		Code: code,
	})

	if err != nil {
		log.Fatalf("Can't read template file: %s", err.Error())
	}

	return s.sender.SendEmail(email, buffer.String())
}

func NewSenderService(sender *mailer.EmailSender) *SenderService {
	return &SenderService{sender: sender}
}
