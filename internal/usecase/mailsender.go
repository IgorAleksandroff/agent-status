package usecase

import (
	"net/smtp"

	"github.com/IgorAleksandroff/agent-status/internal/config"
)

type mailSender struct {
	auth           smtp.Auth
	host, from, to string
}

func NewMailSender(cfg config.MessengerConfig) *mailSender {
	// при двух-факторной авторизации необходимо сгенерить токен вместо пароля
	auth := smtp.PlainAuth("", cfg.Username, cfg.Password, cfg.Host)
	return &mailSender{
		auth: auth,
		host: cfg.Host,
		from: cfg.From,
		to:   cfg.To,
	}
}

func (m mailSender) Send(msg string) error {
	return smtp.SendMail(m.host+":587", m.auth, m.from, []string{m.to}, []byte(msg))
}
