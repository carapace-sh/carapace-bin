package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/yarn"
	"github.com/spf13/cobra"
)

var binCmd = &cobra.Command{
	Use:     "bin",
	Short:   "Get the path to a binary script",
	GroupID: "general",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(binCmd).Standalone()

	binCmd.Flags().Bool("json", false, "Format the output as an NDJSON stream")
	binCmd.Flags().BoolP("verbose", "v", false, "Print both the binary name and the locator of the package that provides the binary")
	rootCmd.AddCommand(binCmd)

	carapace.Gen(binCmd).PositionalCompletion(
		yarn.ActionBins(),
	)
}
