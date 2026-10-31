package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fragments_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set an alias for a fragment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fragments_setCmd).Standalone()

	fragmentsCmd.AddCommand(fragments_setCmd)
}
