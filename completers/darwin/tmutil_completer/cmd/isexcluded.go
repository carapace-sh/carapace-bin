package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(isexcludedCmd)
}

var isexcludedCmd = &cobra.Command{
	Use:   "isexcluded",
	Short: "determine if an item is excluded from backups",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(isexcludedCmd).Standalone()

	isexcludedCmd.Flags().BoolP("xml", "X", false, "Print output in XML property list format")

	carapace.Gen(isexcludedCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}