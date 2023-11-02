package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var help_logCmd = &cobra.Command{
	Use:   "log",
	Short: "Show commit history",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_logCmd).Standalone()

	helpCmd.AddCommand(help_logCmd)
}
