package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deps_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove packages from the project's dependencies",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deps_removeCmd).Standalone()

	deps_removeCmd.Flags().BoolP("help", "h", false, "Print help")
	depsCmd.AddCommand(deps_removeCmd)
}
