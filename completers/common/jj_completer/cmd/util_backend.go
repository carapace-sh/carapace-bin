package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var util_backendCmd = &cobra.Command{
	Use:   "backend",
	Short: "Commands relating to the backend used in the current repo",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(util_backendCmd).Standalone()

	util_backendCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	utilCmd.AddCommand(util_backendCmd)
}
