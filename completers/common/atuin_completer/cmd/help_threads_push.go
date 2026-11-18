package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_threads_pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Capture messages an agent hands over, one JSON object per line on stdin, as capture does its transcript's. Messages already captured are skipped",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_threads_pushCmd).Standalone()

	help_threadsCmd.AddCommand(help_threads_pushCmd)
}
