package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_backup_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a backup of the given share",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_backup_createCmd).Standalone()

	share_backup_createCmd.Flags().String("backup-options", "", "Backup driver option key=value pairs (Optional, Default=None).")
	share_backup_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	share_backup_createCmd.Flags().String("description", "", "Optional share backup description.")
	share_backup_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	share_backup_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	share_backup_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	share_backup_createCmd.Flags().String("name", "", "Optional share backup name.")
	share_backup_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	share_backup_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	share_backup_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	share_backup_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	share_backupCmd.AddCommand(share_backup_createCmd)
}
