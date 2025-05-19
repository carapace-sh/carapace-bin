package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var queueClearCmd = &cobra.Command{
	Use:   "queue:clear [--queue [QUEUE]] [--force] [--] [<connection>]",
	Short: "Delete all of the jobs from the specified queue",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(queueClearCmd).Standalone()

	queueClearCmd.Flags().Bool("force", false, "Force the operation to run when in production")
	queueClearCmd.Flags().String("queue", "", "The name of the queue to clear")
	rootCmd.AddCommand(queueClearCmd)
}
