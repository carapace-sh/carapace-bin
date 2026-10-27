package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_gcCmd = &cobra.Command{
	Use:   "gc",
	Short: "Pack refs, expire reflogs, repack, prune, and refresh commit-graph",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_gcCmd).Standalone()

	helpCmd.AddCommand(help_gcCmd)
}
