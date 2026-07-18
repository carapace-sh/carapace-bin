package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var helpSearchCmd = &cobra.Command{
	Use:   "help-search",
	Short: "Search npm help documentation",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(helpSearchCmd).Standalone()
	helpSearchCmd.Flags().BoolP("long", "l", false, "Show extended information in the search output")

	rootCmd.AddCommand(helpSearchCmd)
}
