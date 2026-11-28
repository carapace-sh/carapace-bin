package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var efuseCommonTableCmd = &cobra.Command{
	Use:   "efuse-common-table",
	Short: "Generate C-source for IDF's eFuse fields",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(efuseCommonTableCmd).Standalone()
	rootCmd.AddCommand(efuseCommonTableCmd)
}
