// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var shellAlias_lsCmd = &cobra.Command{
	Use:     "ls",
	Short:   "List shell aliases",
	Aliases: []string{"list"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(shellAlias_lsCmd).Standalone()

	shellAlias_lsCmd.Flags().BoolP("help", "h", false, "Print help")
	shellAlias_lsCmd.Flags().Bool("no-header", false, "Do not print the table header")
	shellAliasCmd.AddCommand(shellAlias_lsCmd)
}
