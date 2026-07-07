package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var variantCmd = &cobra.Command{
	Use:   "variant",
	Short: "Print the launchd variant",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(variantCmd).Standalone()
	rootCmd.AddCommand(variantCmd)
}
