package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var shimCmd = &cobra.Command{
	Use:   "shim",
	Short: "manage shims",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(shimCmd).Standalone()
	rootCmd.AddCommand(shimCmd)
}
