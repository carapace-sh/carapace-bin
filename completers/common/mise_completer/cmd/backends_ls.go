// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backends_lsCmd = &cobra.Command{
	Use:     "ls",
	Short:   "List built-in backends",
	Aliases: []string{"list"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backends_lsCmd).Standalone()

	backends_lsCmd.Flags().BoolP("help", "h", false, "Print help")
	backendsCmd.AddCommand(backends_lsCmd)
}
