package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var help_chmodCmd = &cobra.Command{
	Use:   "chmod",
	Short: "Sets or removes the executable bit for paths in the repo",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_chmodCmd).Standalone()

	helpCmd.AddCommand(help_chmodCmd)
}
