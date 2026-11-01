package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var whoamiCmd = &cobra.Command{
	Use:   "whoami",
	Short: "Displays your pnpm username",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(whoamiCmd).Standalone()

	whoamiCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(whoamiCmd)
}
