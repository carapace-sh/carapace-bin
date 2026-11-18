package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspaceCmd = &cobra.Command{
	Use:   "workspace",
	Short: "Set up the .atuin.lua config file of a project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaceCmd).Standalone()

	workspaceCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(workspaceCmd)
}
