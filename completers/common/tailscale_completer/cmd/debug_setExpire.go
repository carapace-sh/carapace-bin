package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debug_setExpireCmd = &cobra.Command{
	Use:   "set-expire",
	Short: "Manipulate node key expiry for testing",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_setExpireCmd).Standalone()

	debug_setExpireCmd.Flags().String("in", "", "duration from now to set expiry to")
	debugCmd.AddCommand(debug_setExpireCmd)
}
