package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_catCmd = &cobra.Command{
	Use:   "cat",
	Short: "Show blob, tree, commit, or tag contents for a revision spec",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_catCmd).Standalone()

	helpCmd.AddCommand(help_catCmd)
}
