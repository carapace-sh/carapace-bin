package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resolve_conflictsCmd = &cobra.Command{
	Use:   "conflicts",
	Short: "List the conflicts of a conflicted commit, without entering resolution mode",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resolve_conflictsCmd).Standalone()

	resolve_conflictsCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	resolveCmd.AddCommand(resolve_conflictsCmd)
}
