package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_gc_midxCmd = &cobra.Command{
	Use:   "midx",
	Short: "Write or refresh the multi-pack-index",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_gc_midxCmd).Standalone()

	help_gcCmd.AddCommand(help_gc_midxCmd)
}
