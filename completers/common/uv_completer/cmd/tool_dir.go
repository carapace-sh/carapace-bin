package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tool_dirCmd = &cobra.Command{
	Use:   "dir",
	Short: "Show the path to the uv tools directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tool_dirCmd).Standalone()

	tool_dirCmd.Flags().Bool("bin", false, "Show the directory into which `uv tool` will install executables.")
	toolCmd.AddCommand(tool_dirCmd)
}
