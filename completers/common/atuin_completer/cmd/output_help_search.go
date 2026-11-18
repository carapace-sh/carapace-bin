package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var output_help_searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Full-text search over captured command output",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(output_help_searchCmd).Standalone()

	output_helpCmd.AddCommand(output_help_searchCmd)
}
