package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_environmentCmd = &cobra.Command{
	Use:   "environment",
	Short: "Commands to manage workspace environments",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_environmentCmd).Standalone()

	workspaceCmd.AddCommand(workspace_environmentCmd)
}
