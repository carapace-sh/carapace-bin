package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var python_updateShellCmd = &cobra.Command{
	Use:     "update-shell",
	Short:   "Ensure that the Python executable directory is on the `PATH`",
	Aliases: []string{"ensurepath"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(python_updateShellCmd).Standalone()

	pythonCmd.AddCommand(python_updateShellCmd)
}
