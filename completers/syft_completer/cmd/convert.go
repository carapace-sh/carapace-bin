package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var convertCmd = &cobra.Command{
	Use:   "convert [SOURCE-SBOM] -o [FORMAT]",
	Short: "Convert between SBOM formats",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(convertCmd).Standalone()
	addCommonFlags(convertCmd)
	rootCmd.AddCommand(convertCmd)

	carapace.Gen(convertCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
