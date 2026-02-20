package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var global_help_syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Sync global manifest with installed environments",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(global_help_syncCmd).Standalone()

	global_helpCmd.AddCommand(global_help_syncCmd)
}
