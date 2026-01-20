package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resolve_finishCmd = &cobra.Command{
	Use:   "finish",
	Short: "Finalize conflict resolution and return to workspace mode",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resolve_finishCmd).Standalone()

	resolve_finishCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	resolveCmd.AddCommand(resolve_finishCmd)
}
