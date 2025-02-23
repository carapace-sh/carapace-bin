package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert a recording into another format",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(convertCmd).Standalone()

	convertCmd.Flags().StringP("format", "f", "", "Output file format")
	convertCmd.Flags().BoolP("help", "h", false, "Print help")
	convertCmd.Flags().Bool("overwrite", false, "Overwrite target file if it already exists")
	convertCmd.Flags().BoolP("quiet", "q", false, "Quiet mode, i.e. suppress diagnostic messages")
	convertCmd.Flags().String("server-url", "", "asciinema server URL")
	rootCmd.AddCommand(convertCmd)

	carapace.Gen(convertCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("asciicast", "raw", "txt"),
	})

	carapace.Gen(convertCmd).PositionalCompletion(
		carapace.ActionFiles(),
		carapace.ActionFiles(),
	)
}
