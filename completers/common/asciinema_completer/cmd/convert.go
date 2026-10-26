package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert a recording to another format",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(convertCmd).Standalone()

	convertCmd.Flags().BoolP("help", "h", false, "Print help")
	convertCmd.Flags().StringP("output-format", "f", "", "Output file format")
	convertCmd.Flags().Bool("overwrite", false, "Overwrite the output file if it already exists")
	rootCmd.AddCommand(convertCmd)

	carapace.Gen(convertCmd).FlagCompletion(carapace.ActionMap{
		"output-format": carapace.ActionValues("asciicast-v2", "asciicast-v3", "raw", "txt"),
	})

	carapace.Gen(convertCmd).PositionalCompletion(
		carapace.Batch(
			carapace.ActionFiles(),
			carapace.ActionValuesDescribed("-", "standard input"),
		).ToA(),
		carapace.Batch(
			carapace.ActionFiles(),
			carapace.ActionValuesDescribed("-", "standard output"),
		).ToA(),
	)
}
