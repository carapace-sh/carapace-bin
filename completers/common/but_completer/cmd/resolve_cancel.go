package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resolve_cancelCmd = &cobra.Command{
	Use:   "cancel",
	Short: "Cancel conflict resolution and return to workspace mode",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resolve_cancelCmd).Standalone()

	resolve_cancelCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	resolveCmd.AddCommand(resolve_cancelCmd)
}
