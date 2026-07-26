package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_register_listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List the registered workspaces",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_register_listCmd).Standalone()

	workspace_register_listCmd.Flags().Bool("json", false, "Output in JSON format")
	workspace_registerCmd.AddCommand(workspace_register_listCmd)
}
