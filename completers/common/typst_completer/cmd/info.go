package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Displays debugging information about Typst",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(infoCmd).Standalone()

	infoCmd.Flags().StringP("format", "f", "", "The format to serialize in, if it should be machine-readable")
	infoCmd.Flags().Bool("pretty", false, "Whether to pretty-print the serialized output")
	rootCmd.AddCommand(infoCmd)

	carapace.Gen(infoCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "yaml"),
	})
}
