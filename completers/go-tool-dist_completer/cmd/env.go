package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var envCmd = &cobra.Command{
	Use:   "env",
	Short: "print environment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(envCmd).Standalone()

	envCmd.Flags().BoolS("9", "9", false, "emit plan 9 syntax")
	envCmd.Flags().BoolS("p", "p", false, "emit updated $PATH")
	envCmd.Flags().BoolS("v", "v", false, "verbosity")
	envCmd.Flags().BoolS("w", "w", false, "emit windows syntax")

	rootCmd.AddCommand(envCmd)
}
