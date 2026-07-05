package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var productCmd = &cobra.Command{
	Use:   "product",
	Short: "installed products",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(productCmd).Standalone()
	rootCmd.AddCommand(productCmd)
}
