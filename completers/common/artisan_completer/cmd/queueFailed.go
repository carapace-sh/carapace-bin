package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var queueFailedCmd = &cobra.Command{
	Use:   "queue:failed",
	Short: "List all of the failed queue jobs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(queueFailedCmd).Standalone()

	rootCmd.AddCommand(queueFailedCmd)
}
