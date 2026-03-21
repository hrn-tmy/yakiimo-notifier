package email

import (
	"fmt"
	"net/smtp"
	"strings"
)

type SMTPSender struct {
	host string
	port string
	from string
}

func NewSMTPSender(host, port, from string) *SMTPSender {
	return &SMTPSender{host: host, port: port, from: from}
}

func (s *SMTPSender) Send(to []string, subject, body string) error {
	addr := fmt.Sprintf("%s:%s", s.host, s.port)
	msg := fmt.Sprintf(
		"From: %s\r\nTo: %s\r\nSubject: %s\r\nContent-Type: text/plain; charset=UTF-8\r\n\r\n%s",
		s.from,
		strings.Join(to, ", "),
		subject,
		body,
	)
	return smtp.SendMail(addr, nil, s.from, to, []byte(msg))
}
