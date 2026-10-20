// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var installs_selectCmd = &cobra.Command{
	Use:   "select",
	Short: "Choose which installation mise uses when no lockfile pins one",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(installs_selectCmd).Standalone()

	installs_selectCmd.Flags().BoolP("help", "h", false, "Print help")
	installsCmd.AddCommand(installs_selectCmd)
}
