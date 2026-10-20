// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var direnv_activateCmd = &cobra.Command{
	Use:    "activate",
	Short:  "Print `use_mise`, a direnv function that loads mise inside direnv",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(direnv_activateCmd).Standalone()

	direnv_activateCmd.Flags().BoolP("help", "h", false, "Print help")
	direnvCmd.AddCommand(direnv_activateCmd)
}
