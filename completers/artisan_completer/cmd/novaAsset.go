package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var novaAssetCmd = &cobra.Command{
	Use:   "nova:asset <name>",
	Short: "Create a new asset",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(novaAssetCmd).Standalone()

	rootCmd.AddCommand(novaAssetCmd)
}
