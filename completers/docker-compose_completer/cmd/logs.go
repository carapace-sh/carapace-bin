package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/docker-compose_completer/cmd/action"
	"github.com/spf13/cobra"
)

var logsCmd = &cobra.Command{
	Use:   "logs",
	Short: "View output from containers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logsCmd).Standalone()

	logsCmd.Flags().BoolP("follow", "f", false, "Follow log output.")
	logsCmd.Flags().Bool("no-color", false, "Produce monochrome output.")
	logsCmd.Flags().String("tail", "", "Number of lines to show from the end of the logs")
	logsCmd.Flags().BoolP("timestamps", "t", false, "Show timestamps.")
	rootCmd.AddCommand(logsCmd)

	carapace.Gen(logsCmd).PositionalAnyCompletion(
		action.ActionServices(logsCmd),
	)
}
