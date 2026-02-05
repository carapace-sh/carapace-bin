package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_rootCmd = &cobra.Command{
	Use:   "root",
	Short: "Show the workspace root directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_rootCmd).Standalone()

	workspace_rootCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	workspace_rootCmd.Flags().String("name", "", "Name of the workspace (defaults to current)")
	workspaceCmd.AddCommand(workspace_rootCmd)
}
