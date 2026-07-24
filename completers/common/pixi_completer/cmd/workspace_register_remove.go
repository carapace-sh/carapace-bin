package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_register_removeCmd = &cobra.Command{
	Use:     "remove",
	Short:   "Remove a workspace from registry",
	Aliases: []string{"rm"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_register_removeCmd).Standalone()

	workspace_registerCmd.AddCommand(workspace_register_removeCmd)
}
