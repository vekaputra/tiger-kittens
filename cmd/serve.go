package cmd

import (
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

func newServeCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "serve",
		Short: "Run HTTP server",
		Run:   func(cmd *cobra.Command, args []string) {},
	}
}
