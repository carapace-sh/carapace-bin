package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var volume_summaryCmd = &cobra.Command{
	Use:   "summary",
	Short: "Show a summary of all volumes in this deployment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volume_summaryCmd).Standalone()

	volume_summaryCmd.Flags().Bool("all-projects", false, "Include all projects (admin only)")
	volume_summaryCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	volume_summaryCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	volume_summaryCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	volume_summaryCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	volume_summaryCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	volume_summaryCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	volume_summaryCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	volume_summaryCmd.Flags().String("variable", "", "==SUPPRESS==")
	volumeCmd.AddCommand(volume_summaryCmd)
}
