package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_viewCmd = &cobra.Command{
	Use:   "view",
	Short: "View registry information about a package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_viewCmd).Standalone()

	helpCmd.AddCommand(help_viewCmd)
}
