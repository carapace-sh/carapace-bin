package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesCmd = &cobra.Command{
	Use:     "workspaces",
	Short:   "manage workspaces",
	GroupID: "workspace-related",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesCmd).Standalone()

	rootCmd.AddCommand(workspacesCmd)
}
