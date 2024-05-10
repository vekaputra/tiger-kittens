package mailqueue

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/vekaputra/tiger-kittens/internal/repository/pgsql"
)

//go:generate mockery --name=Provider --outpkg=mock --output=./mock
type Provider interface {
	Add(job SendEmailJob)
}

type MailQueue struct {
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

func New(tigerRepository pgsql.TigerRepositoryProvider, userRepository pgsql.UserRepositoryProvider) *MailQueue {
	return &MailQueue{
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
			log.Info().Msg(fmt.Sprintf("send sighting email to: %s", job.DestinationEmail))
			continue
		}
	}
}

func (q *MailQueue) Add(job SendEmailJob) {
	q.jobs <- job
	log.Info().Msg(fmt.Sprintf("new send email job added for: %s", job.DestinationEmail))
}
