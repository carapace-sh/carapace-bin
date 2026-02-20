package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search a conda package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_searchCmd).Standalone()

	helpCmd.AddCommand(help_searchCmd)
}
