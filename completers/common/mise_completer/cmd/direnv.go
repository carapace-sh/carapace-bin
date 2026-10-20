// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var direnvCmd = &cobra.Command{
	Use:    "direnv",
	Short:  "Print `use_mise`, a direnv function that loads mise inside direnv",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(direnvCmd).Standalone()

	direnvCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(direnvCmd)
}
