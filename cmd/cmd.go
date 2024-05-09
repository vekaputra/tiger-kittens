package cmd

import "github.com/spf13/cobra"

type Config struct {
	// Use is the one-line usage message.
	Use string

	// Short is the short description shown in the 'help' output.
	Short string

	// Long is the long message shown in the 'help <this-command>' output.
	Long string
}

func New(c Config, db DB) *cobra.Command {
	cmd := &cobra.Command{
		Use:   c.Use,
		Short: c.Short,
		Long:  c.Long,
	}

	cmd.AddCommand(newServeCommand())
	cmd.AddCommand(newDBMigrateCommand(db))
	cmd.AddCommand(newDBRollbackCommand(db))

	return cmd
}
