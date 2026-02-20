package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_globalCmd = &cobra.Command{
	Use:   "global",
	Short: "Subcommand for global package management actions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_globalCmd).Standalone()

	helpCmd.AddCommand(help_globalCmd)
}
