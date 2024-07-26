package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var queueRestartCmd = &cobra.Command{
	Use:   "queue:restart",
	Short: "Restart queue worker daemons after their current job",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(queueRestartCmd).Standalone()

	rootCmd.AddCommand(queueRestartCmd)
}
