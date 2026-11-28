package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var efuseCustomTableCmd = &cobra.Command{
	Use:   "efuse-custom-table",
	Short: "Generate C-source for user's eFuse fields",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(efuseCustomTableCmd).Standalone()
	rootCmd.AddCommand(efuseCustomTableCmd)
}
