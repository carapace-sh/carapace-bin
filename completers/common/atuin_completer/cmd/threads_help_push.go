package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var threads_help_pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Capture messages an agent hands over, one JSON object per line on stdin, as capture does its transcript's. Messages already captured are skipped",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(threads_help_pushCmd).Standalone()

	threads_helpCmd.AddCommand(threads_help_pushCmd)
}
