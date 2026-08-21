package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var volume_backup_record_exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export volume backup details.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volume_backup_record_exportCmd).Standalone()

	volume_backup_record_exportCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	volume_backup_record_exportCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	volume_backup_record_exportCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	volume_backup_record_exportCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	volume_backup_record_exportCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	volume_backup_record_exportCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	volume_backup_record_exportCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	volume_backup_record_exportCmd.Flags().String("variable", "", "==SUPPRESS==")
	volume_backup_recordCmd.AddCommand(volume_backup_record_exportCmd)
}
