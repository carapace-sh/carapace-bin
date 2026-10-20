// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var daemonsCmd = &cobra.Command{
	Use:     "daemons",
	Short:   "[experimental] Manage project daemons with pitchfork",
	Aliases: []string{"daemon"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(daemonsCmd).Standalone()

	daemonsCmd.Flags().BoolP("help", "h", false, "Print help")
	daemonsCmd.Flags().Bool("json", false, "Output in JSON format, including ports, URLs, and data directories")
	rootCmd.AddCommand(daemonsCmd)
}
