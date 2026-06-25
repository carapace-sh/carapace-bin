package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var variantCmd = &cobra.Command{
	Use:   "variant",
	Short: "Prints the launchd variant",
	Run:   func(*cobra.Command, []string) {},
}

func init() {
	carapace.Gen(variantCmd).Standalone()
	rootCmd.AddCommand(variantCmd)
}
