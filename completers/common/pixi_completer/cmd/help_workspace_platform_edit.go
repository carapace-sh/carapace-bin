package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_workspace_platform_editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit an existing workspace platform's subdir and/or virtual packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_workspace_platform_editCmd).Standalone()

	help_workspace_platformCmd.AddCommand(help_workspace_platform_editCmd)
}
