package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var volume_attachment_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show detailed information for a volume attachment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volume_attachment_showCmd).Standalone()

	volume_attachment_showCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	volume_attachment_showCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	volume_attachment_showCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	volume_attachment_showCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	volume_attachment_showCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	volume_attachment_showCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	volume_attachment_showCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	volume_attachment_showCmd.Flags().String("variable", "", "==SUPPRESS==")
	volume_attachmentCmd.AddCommand(volume_attachment_showCmd)
}
