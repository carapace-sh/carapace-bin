package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var help_utilCmd = &cobra.Command{
	Use:   "util",
	Short: "Infrequently used commands such as for generating shell completions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_utilCmd).Standalone()

	helpCmd.AddCommand(help_utilCmd)
}
