package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_lintersCmd = &cobra.Command{
	Use:   "linters",
	Short: "Display help for linters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_lintersCmd).Standalone()

	help_lintersCmd.Flags().Bool("json", false, "Display as JSON")
	helpCmd.AddCommand(help_lintersCmd)
}
