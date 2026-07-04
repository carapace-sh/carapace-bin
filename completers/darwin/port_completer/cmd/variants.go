package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var variantsCmd = &cobra.Command{
	Use:   "variants",
	Short: "List available variants for a port",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(variantsCmd).Standalone()
	rootCmd.AddCommand(variantsCmd)
}
