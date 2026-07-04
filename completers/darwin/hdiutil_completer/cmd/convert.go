package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert a disk image from one format to another",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(convertCmd).Standalone()

	convertCmd.Flags().String("format", "", "Target format: UDZO, UDCO, UDRO, UDRW, UDSP, UDBZ, UDIF")
	convertCmd.Flags().StringP("output", "o", "", "Output file path")
	convertCmd.Flags().Bool("plist", false, "Provide result output in plist format")
	convertCmd.Flags().Bool("quiet", false, "Close stdout and stderr")
	convertCmd.Flags().Bool("verbose", false, "Produce extra progress output")
	rootCmd.AddCommand(convertCmd)

	carapace.Gen(convertCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("UDZO", "UDCO", "UDRO", "UDRW", "UDSP", "UDBZ", "UDIF"),
		"output": carapace.ActionFiles(),
	})

	carapace.Gen(convertCmd).PositionalCompletion(carapace.ActionFiles())
}
