// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sync_pythonCmd = &cobra.Command{
	Use:   "python",
	Short: "Symlink python versions installed by pyenv or uv into mise",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sync_pythonCmd).Standalone()

	sync_pythonCmd.Flags().BoolP("help", "h", false, "Print help")
	sync_pythonCmd.Flags().Bool("pyenv", false, "Get tool versions from pyenv")
	sync_pythonCmd.Flags().Bool("uv", false, "Link uv's python versions into mise and mise's python versions into uv")
	syncCmd.AddCommand(sync_pythonCmd)
}
