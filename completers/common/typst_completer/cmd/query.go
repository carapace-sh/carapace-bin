package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/typst_completer/cmd/common"
	"github.com/spf13/cobra"
)

var queryCmd = &cobra.Command{
	Use:   "query",
	Short: "Processes an input file to extract provided metadata",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(queryCmd).Standalone()

	queryCmd.Flags().BoolP("help", "h", false, "Print help (see a summary with '-h')")

	common.AddBuildFlags(queryCmd)

	rootCmd.AddCommand(queryCmd)

	// TODO: PositionalCompletion for templates
}
