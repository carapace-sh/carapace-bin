package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var installScripts_pruneCmd = &cobra.Command{
	Use:   "prune",
	Short: "Drop allowScripts entries that no longer match an installed package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(installScripts_pruneCmd).Standalone()

	installScriptsCmd.AddCommand(installScripts_pruneCmd)
}
