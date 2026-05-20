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

	queryCmd.Flags().String("field", "", "Extracts just one field from all retrieved elements")
	queryCmd.Flags().BoolP("help", "h", false, "Print help (see a summary with '-h')")
	queryCmd.Flags().Bool("one", false, "Expects and retrieves exactly one element")
	queryCmd.Flags().Bool("pretty", false, "Whether to pretty-print the serialized output")
	queryCmd.Flags().String("target", "", "The target to compile for")
	common.AddBuildFlags(queryCmd)
	rootCmd.AddCommand(queryCmd)

	carapace.Gen(queryCmd).FlagCompletion(carapace.ActionMap{
		"target": carapace.ActionValuesDescribed(
			"paged", "PDF and image formats",
			"html", "HTML",
		),
	})

	// TODO: PositionalCompletion for templates
}
