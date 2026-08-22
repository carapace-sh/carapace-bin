package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var shimCmd = &cobra.Command{
	Use:   "shim",
	Short: "manipulate Scoop shims",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(shimCmd).Standalone()
	shimCmd.Flags().BoolP("global", "g", false, "manipulate global shim(s)")
	rootCmd.AddCommand(shimCmd)
}
