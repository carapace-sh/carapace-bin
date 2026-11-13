package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fragments_loadersCmd = &cobra.Command{
	Use:   "loaders",
	Short: "Show fragment loaders registered by plugins",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fragments_loadersCmd).Standalone()

	fragmentsCmd.AddCommand(fragments_loadersCmd)
}
