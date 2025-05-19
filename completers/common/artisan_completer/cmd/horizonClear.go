package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var horizonClearCmd = &cobra.Command{
	Use:   "horizon:clear [--queue [QUEUE]] [--force]",
	Short: "Delete all of the jobs from the specified queue",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(horizonClearCmd).Standalone()

	horizonClearCmd.Flags().Bool("force", false, "Force the operation to run when in production")
	horizonClearCmd.Flags().String("queue", "", "The name of the queue to clear")
	rootCmd.AddCommand(horizonClearCmd)
}
