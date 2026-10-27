package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_ignoreCmd = &cobra.Command{
	Use:   "ignore",
	Short: "Query exclude (gitignore) rules",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_ignoreCmd).Standalone()

	helpCmd.AddCommand(help_ignoreCmd)
}
