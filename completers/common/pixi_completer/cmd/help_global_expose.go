package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_global_exposeCmd = &cobra.Command{
	Use:   "expose",
	Short: "Interact with the exposure of binaries in the global environment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_global_exposeCmd).Standalone()

	help_globalCmd.AddCommand(help_global_exposeCmd)
}
