// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var daemons_stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "[experimental] Stop project daemons",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(daemons_stopCmd).Standalone()

	daemons_stopCmd.Flags().BoolP("help", "h", false, "Print help")
	daemonsCmd.AddCommand(daemons_stopCmd)
}
