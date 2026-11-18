package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create a .atuin.lua workspace config file in the current directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_initCmd).Standalone()

	workspace_initCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	workspace_initCmd.Flags().StringSlice("org", nil, "An organization to share the threads with, by slug. Repeat it to name more than one. Without it, you're asked.")
	workspaceCmd.AddCommand(workspace_initCmd)
}
