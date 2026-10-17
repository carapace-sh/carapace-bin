package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var blockTypesCmd = &cobra.Command{
	Use:   "types",
	Short: "List block types",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(blockTypesCmd).Standalone()

	blockTypesCmd.Flags().Bool("json", false, "output block types as JSON")

	blockCmd.AddCommand(blockTypesCmd)
}
