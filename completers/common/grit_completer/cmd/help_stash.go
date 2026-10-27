package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_stashCmd = &cobra.Command{
	Use:   "stash",
	Short: "Save or restore shelved work",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_stashCmd).Standalone()

	helpCmd.AddCommand(help_stashCmd)
}
