package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debug_rebindCmd = &cobra.Command{
	Use:   "rebind",
	Short: "Force a magicsock rebind",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_rebindCmd).Standalone()

	debugCmd.AddCommand(debug_rebindCmd)
}
