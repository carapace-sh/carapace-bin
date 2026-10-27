package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gc_help_midxCmd = &cobra.Command{
	Use:   "midx",
	Short: "Write or refresh the multi-pack-index",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gc_help_midxCmd).Standalone()

	gc_helpCmd.AddCommand(gc_help_midxCmd)
}
