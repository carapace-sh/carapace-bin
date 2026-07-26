package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_help_platform_editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit an existing workspace platform's subdir and/or virtual packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_help_platform_editCmd).Standalone()

	workspace_help_platformCmd.AddCommand(workspace_help_platform_editCmd)
}
