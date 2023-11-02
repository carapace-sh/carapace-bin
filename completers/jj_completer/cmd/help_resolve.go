package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var help_resolveCmd = &cobra.Command{
	Use:   "resolve",
	Short: "Resolve a conflicted file with an external merge tool",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_resolveCmd).Standalone()

	helpCmd.AddCommand(help_resolveCmd)
}
