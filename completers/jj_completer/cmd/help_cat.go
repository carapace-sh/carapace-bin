package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var help_catCmd = &cobra.Command{
	Use:   "cat",
	Short: "Print contents of a file in a revision",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_catCmd).Standalone()

	helpCmd.AddCommand(help_catCmd)
}
