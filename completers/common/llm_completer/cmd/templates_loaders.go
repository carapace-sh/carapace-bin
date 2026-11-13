package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var templates_loadersCmd = &cobra.Command{
	Use:   "loaders",
	Short: "Show template loaders registered by plugins",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(templates_loadersCmd).Standalone()

	templatesCmd.AddCommand(templates_loadersCmd)
}
