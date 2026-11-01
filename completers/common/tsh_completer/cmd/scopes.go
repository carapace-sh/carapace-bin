package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var scopesCmd = &cobra.Command{
	Use:     "scopes",
	Short:   "View and manage Teleport scopes.",
	Aliases: []string{"scoped"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(scopesCmd).Standalone()

	rootCmd.AddCommand(scopesCmd)
}
