package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_repackCmd = &cobra.Command{
	Use:   "repack",
	Short: "Pack reachable objects and remove redundant packs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_repackCmd).Standalone()

	helpCmd.AddCommand(help_repackCmd)
}
