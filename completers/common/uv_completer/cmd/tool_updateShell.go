package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tool_updateShellCmd = &cobra.Command{
	Use:     "update-shell",
	Short:   "Ensure that the tool executable directory is on the `PATH`",
	Aliases: []string{"ensurepath"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tool_updateShellCmd).Standalone()

	toolCmd.AddCommand(tool_updateShellCmd)
}
