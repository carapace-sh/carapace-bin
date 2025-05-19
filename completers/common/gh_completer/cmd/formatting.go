package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var formattingCmd = &cobra.Command{
	Use:    "formatting",
	Short:  "Formatting options for JSON data exported from gh",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(formattingCmd).Standalone()

	rootCmd.AddCommand(formattingCmd)
}
