package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

type Server interface {
	// GracefulStop stops the server gracefully. It stops the server from accepting new connections
	// and blocks until all the pending requests are finished.
	GracefulStop() error

	// Serve accepts incoming connections on the listener lis, creating a new ServerTransport and
	// service goroutine for each. The service goroutines read gRPC requests and then call the
	// registered handlers to reply to them. Serve returns when lis.Accept fails with fatal errors.
	// lis will be closed when this method returns. Serve will return a non-nil error unless Stop
	// or GracefulStop is called.
	Serve() error
}

type Client interface {
	// Close closes the client, releasing any open resources.
	Close() error
}

func newServeCommand(srv Server, dbClient Client) *cobra.Command {
	return &cobra.Command{
		Use:   "serve",
		Short: "Run HTTP server",
		Run: func(cmd *cobra.Command, args []string) {
			go func() {
				if err := srv.Serve(); err != nil {
					log.Info().Msg(fmt.Sprintf("echo_server has shut down: %v", err))
				}
			}()

			waitForShutdown()

			log.Info().Msg("shutting down the server")
			if err := srv.GracefulStop(); err != nil {
				log.Error().Err(err).Send()
			}

			if err := dbClient.Close(); err != nil {
				log.Error().Err(err).Send()
			}
		},
	}
}

func waitForShutdown() os.Signal { //nolint:unparam
	ch := make(chan os.Signal, 4)
	signal.Notify(
		ch,
		syscall.SIGINT,
		syscall.SIGQUIT,
		syscall.SIGTERM,
		syscall.SIGTSTP,
	)
	return <-ch
}
