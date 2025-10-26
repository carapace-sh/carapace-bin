package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Extract a value from the given TOML document",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(getCmd).Standalone()

	getCmd.Flags().StringP("file-path", "f", "", "Path to the TOML document, if omitted the standard input will be used")
	getCmd.Flags().StringP("output-format", "o", "", "The format specifying how the output is printed")
	getCmd.Flags().String("separator", "", "A string that separates array values when printing to stdout")
	getCmd.Flags().BoolP("strip-newline", "s", false, "Strip the trailing newline from the output")
	rootCmd.AddCommand(getCmd)

	carapace.Gen(getCmd).FlagCompletion(carapace.ActionMap{
		"file-path": carapace.ActionFiles(),
		"output-format": carapace.ActionValuesDescribed(
			"value", "Extract the value outputting it in a text format",
			"json", "Output format in JSON",
			"toml", "Output format in TOML",
		),
	})
}
