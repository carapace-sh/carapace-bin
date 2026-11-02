package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fragments_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a fragment alias",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fragments_removeCmd).Standalone()

	fragmentsCmd.AddCommand(fragments_removeCmd)
}
