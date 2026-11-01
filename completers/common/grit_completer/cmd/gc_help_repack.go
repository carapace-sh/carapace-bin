package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gc_help_repackCmd = &cobra.Command{
	Use:   "repack",
	Short: "Repack reachable objects into one pack",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gc_help_repackCmd).Standalone()

	gc_helpCmd.AddCommand(gc_help_repackCmd)
}
