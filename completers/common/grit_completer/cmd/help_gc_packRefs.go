package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_gc_packRefsCmd = &cobra.Command{
	Use:   "pack-refs",
	Short: "Pack loose refs into `packed-refs` (or compact reftable)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_gc_packRefsCmd).Standalone()

	help_gcCmd.AddCommand(help_gc_packRefsCmd)
}
