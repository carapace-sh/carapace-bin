package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bgp_speaker_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show a BGP speaker",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bgp_speaker_showCmd).Standalone()

	bgp_speaker_showCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	bgp_speaker_showCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	bgp_speaker_showCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	bgp_speaker_showCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	bgp_speaker_showCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	bgp_speaker_showCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	bgp_speaker_showCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	bgp_speaker_showCmd.Flags().String("variable", "", "==SUPPRESS==")
	bgp_speakerCmd.AddCommand(bgp_speaker_showCmd)
}
