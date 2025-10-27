package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_formattersCmd = &cobra.Command{
	Use:   "formatters",
	Short: "Display help for formatters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_formattersCmd).Standalone()

	help_formattersCmd.Flags().Bool("json", false, "Display as JSON")
	helpCmd.AddCommand(help_formattersCmd)
}
