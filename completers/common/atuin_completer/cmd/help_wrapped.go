package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_wrappedCmd = &cobra.Command{
	Use:   "wrapped",
	Short: "Show a fun, year-in-review recap of your shell history",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_wrappedCmd).Standalone()

	helpCmd.AddCommand(help_wrappedCmd)
}
