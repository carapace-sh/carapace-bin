package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_global_syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Sync global manifest with installed environments",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_global_syncCmd).Standalone()

	help_globalCmd.AddCommand(help_global_syncCmd)
}
