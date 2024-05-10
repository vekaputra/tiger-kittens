package mailqueue

import (
	"context"
	"fmt"
	"net/smtp"

	"github.com/rs/zerolog/log"
	"github.com/vekaputra/tiger-kittens/internal/config"
	"github.com/vekaputra/tiger-kittens/internal/repository/pgsql"
)

//go:generate mockery --name=Provider --outpkg=mock --output=./mock
type Provider interface {
	Add(job SendEmailJob)
}

type Config struct {
	AppPassword     string
	EmailEnabled    bool
	EmailServerAddr string
	EmailServerHost string
	SenderEmail     string
}

type MailQueue struct {
	config          Config
	cancel          context.CancelFunc
	ctx             context.Context
	jobs            chan SendEmailJob
	tigerRepository pgsql.TigerRepositoryProvider
	userRepository  pgsql.UserRepositoryProvider
}

type SendEmailJob struct {
	DestinationEmail string
	Title            string
	Body             string
}

func New(config config.EmailConfig, tigerRepository pgsql.TigerRepositoryProvider, userRepository pgsql.UserRepositoryProvider) *MailQueue {
	return &MailQueue{
		config: Config{
			AppPassword:     config.AppPassword,
			EmailServerAddr: config.EmailServerAddr,
			EmailServerHost: config.EmailServerHost,
			EmailEnabled:    config.EmailEnabled,
			SenderEmail:     config.SenderEmail,
		},
		tigerRepository: tigerRepository,
		userRepository:  userRepository,
	}
}

func (q *MailQueue) GracefulStop() error {
	q.cancel()
	return nil
}

func (q *MailQueue) Serve() error {
	log.Info().Msg(fmt.Sprintf("starting send email queue worker"))

	ctx, cancel := context.WithCancel(context.Background())
	q.ctx = ctx
	q.cancel = cancel
	q.jobs = make(chan SendEmailJob)

	for {
		select {
		case <-q.ctx.Done():
			log.Info().Msg("context canceled, stopping worker")
			return nil
		case job := <-q.jobs:
			if err := q.sendMail(job); err != nil {
				log.Error().Err(err).Msg(fmt.Sprintf("failed to send email to: %s", job.DestinationEmail))
				continue
			}
			log.Info().Msg(fmt.Sprintf("send sighting email to: %s", job.DestinationEmail))
		}
	}
}

func (q *MailQueue) Add(job SendEmailJob) {
	q.jobs <- job
	log.Info().Msg(fmt.Sprintf("new send email job added for: %s", job.DestinationEmail))
}

func (q *MailQueue) sendMail(job SendEmailJob) error {
	if !q.config.EmailEnabled {
		return nil
	}

	auth := smtp.PlainAuth("", q.config.SenderEmail, q.config.AppPassword, q.config.EmailServerHost)
	to := []string{job.DestinationEmail}
	msg := []byte(fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s", job.DestinationEmail, job.Title, job.Body))
	return smtp.SendMail(q.config.EmailServerAddr, auth, q.config.SenderEmail, to, msg)
}
