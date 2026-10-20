// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var daemons_registerCmd = &cobra.Command{
	Use:   "register",
	Short: "[experimental] Register project daemons for on-demand startup",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(daemons_registerCmd).Standalone()

	daemons_registerCmd.Flags().BoolP("help", "h", false, "Print help")
	daemonsCmd.AddCommand(daemons_registerCmd)
}
