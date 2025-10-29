package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_manCmd = &cobra.Command{
	Use:   "man",
	Short: "View the man page for a given command",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_manCmd).Standalone()

	helpCmd.AddCommand(help_manCmd)
}
