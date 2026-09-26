package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var python_dirCmd = &cobra.Command{
	Use:   "dir",
	Short: "Show the uv Python installation directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(python_dirCmd).Standalone()

	python_dirCmd.Flags().Bool("bin", false, "Show the directory into which `uv python` will install Python executables.")
	pythonCmd.AddCommand(python_dirCmd)
}
