package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var util_gcCmd = &cobra.Command{
	Use:   "gc",
	Short: "Run backend-dependent garbage collection",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(util_gcCmd).Standalone()

	util_gcCmd.Flags().String("expire", "", "Time threshold")
	util_gcCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	utilCmd.AddCommand(util_gcCmd)
}
