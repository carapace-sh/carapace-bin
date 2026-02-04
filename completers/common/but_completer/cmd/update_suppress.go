package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var update_suppressCmd = &cobra.Command{
	Use:   "suppress",
	Short: "Suppress update notifications temporarily",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(update_suppressCmd).Standalone()

	update_suppressCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	updateCmd.AddCommand(update_suppressCmd)
}
