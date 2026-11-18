package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_output_searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Full-text search over captured command output",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_output_searchCmd).Standalone()

	help_outputCmd.AddCommand(help_output_searchCmd)
}
