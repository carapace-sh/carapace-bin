package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gc_midxCmd = &cobra.Command{
	Use:   "midx",
	Short: "Write or refresh the multi-pack-index",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gc_midxCmd).Standalone()

	gc_midxCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	gcCmd.AddCommand(gc_midxCmd)
}
