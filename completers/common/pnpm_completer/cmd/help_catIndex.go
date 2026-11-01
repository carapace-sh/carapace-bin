package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_catIndexCmd = &cobra.Command{
	Use:   "cat-index",
	Short: "Prints the index file of a specific package from the store",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_catIndexCmd).Standalone()

	helpCmd.AddCommand(help_catIndexCmd)
}
