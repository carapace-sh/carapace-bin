package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "textutil",
	Short: "text converter utility",
	Long:  "https://keith.github.io/xcode-manpages/textutil.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("convert", "convert", "", "Convert to the specified format")
	rootCmd.Flags().StringP("encoding", "encoding", "", "Set the encoding")
	rootCmd.Flags().StringP("format", "format", "", "Set the format")
	rootCmd.Flags().BoolP("help", "h", false, "Display usage information")
	rootCmd.Flags().BoolP("info", "info", false, "Print information about the file")
	rootCmd.Flags().StringP("input", "input", "", "Input file")
	rootCmd.Flags().StringP("output", "output", "", "Output file")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"convert": carapace.ActionValues("txt", "html", "rtf", "rtfd", "doc", "wordml", "webarchive"),
		"format":  carapace.ActionValues("txt", "html", "rtf", "rtfd", "doc", "wordml", "webarchive"),
		"input":   carapace.ActionFiles(),
		"output":  carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
