package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var stack_syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Sync and submit progress on a stacked diff. (EXPERIMENTAL)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stack_syncCmd).Standalone()

	stackCmd.AddCommand(stack_syncCmd)
}
