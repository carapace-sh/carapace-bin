// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sponsorsCmd = &cobra.Command{
	Use:   "sponsors",
	Short: "Show the companies sponsoring mise and the jdx.dev open source tools",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sponsorsCmd).Standalone()

	sponsorsCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(sponsorsCmd)
}
