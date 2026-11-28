package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var showEfuseTableCmd = &cobra.Command{
	Use:   "show-efuse-table",
	Short: "Print eFuse table",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(showEfuseTableCmd).Standalone()

	rootCmd.AddCommand(showEfuseTableCmd)
}
