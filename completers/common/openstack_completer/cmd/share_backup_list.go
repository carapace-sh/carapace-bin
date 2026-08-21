package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_backup_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List share backups",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_backup_listCmd).Standalone()

	share_backup_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	share_backup_listCmd.Flags().String("description", "", "Filter results by description.")
	share_backup_listCmd.Flags().String("description~", "", "Filter results matching a share backup description ")
	share_backup_listCmd.Flags().String("detail", "", "Show detailed information about share backups.")
	share_backup_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	share_backup_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	share_backup_listCmd.Flags().String("limit", "", "Limit the number of backups returned.")
	share_backup_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	share_backup_listCmd.Flags().String("name", "", "Filter results by name.")
	share_backup_listCmd.Flags().String("name~", "", "Filter results matching a share backup name pattern. ")
	share_backup_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	share_backup_listCmd.Flags().String("offset", "", "Start position of backup records listing.")
	share_backup_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	share_backup_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	share_backup_listCmd.Flags().String("share", "", "Name or ID of the share to list backups for.")
	share_backup_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	share_backup_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	share_backup_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	share_backup_listCmd.Flags().String("sort_dir", "", "Sort direction, available values are ('asc', 'desc').")
	share_backup_listCmd.Flags().String("sort_key", "", "Key to be sorted, available keys are ('id', 'status', 'size', 'share_id', 'progress', 'restore_progress', 'name', 'host', 'topic', 'project_id').")
	share_backup_listCmd.Flags().String("status", "", "Filter results by status.")
	share_backupCmd.AddCommand(share_backup_listCmd)
}
