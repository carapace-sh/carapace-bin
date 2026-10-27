package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_gc_repackCmd = &cobra.Command{
	Use:   "repack",
	Short: "Repack reachable objects into one pack",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_gc_repackCmd).Standalone()

	help_gcCmd.AddCommand(help_gc_repackCmd)
}
