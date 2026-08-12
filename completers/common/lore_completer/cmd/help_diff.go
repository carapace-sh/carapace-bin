package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_diffCmd = &cobra.Command{
	Use:   "diff",
	Short: "Show differences between two revisions of a file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_diffCmd).Standalone()

	helpCmd.AddCommand(help_diffCmd)
}
