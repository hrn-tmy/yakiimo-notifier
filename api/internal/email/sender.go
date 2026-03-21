package email

type Sender interface {
	Send(to []string, subject, body string) error
}
