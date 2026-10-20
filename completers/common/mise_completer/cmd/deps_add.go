// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deps_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add packages to the project's dependencies",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deps_addCmd).Standalone()

	deps_addCmd.Flags().BoolP("dev", "D", false, "Add as a development dependency")
	deps_addCmd.Flags().BoolP("help", "h", false, "Print help")
	depsCmd.AddCommand(deps_addCmd)
}
