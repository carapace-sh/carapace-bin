package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gc_help_packRefsCmd = &cobra.Command{
	Use:   "pack-refs",
	Short: "Pack loose refs into `packed-refs` (or compact reftable)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gc_help_packRefsCmd).Standalone()

	gc_helpCmd.AddCommand(gc_help_packRefsCmd)
}
