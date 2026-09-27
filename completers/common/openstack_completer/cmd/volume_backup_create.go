package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var volume_backup_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create new volume backup",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volume_backup_createCmd).Standalone()

	volume_backup_createCmd.Flags().String("availability-zone", "", "AZ where the backup should be stored; by default it will be the same as the source (supported by --os-volume-api-version 3.51 or above)")
	volume_backup_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	volume_backup_createCmd.Flags().String("container", "", "Optional backup container name")
	volume_backup_createCmd.Flags().String("description", "", "Description of the backup")
	volume_backup_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	volume_backup_createCmd.Flags().Bool("force", false, "Allow to back up an in-use volume")
	volume_backup_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	volume_backup_createCmd.Flags().Bool("incremental", false, "Perform an incremental backup")
	volume_backup_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	volume_backup_createCmd.Flags().String("name", "", "Name of the backup")
	volume_backup_createCmd.Flags().Bool("no-incremental", false, "Do not perform an incremental backup")
	volume_backup_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	volume_backup_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	volume_backup_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	volume_backup_createCmd.Flags().String("property", "", "Set a property on this backup (repeat option to remove multiple values) (supported by --os-volume-api-version 3.43 or above)")
	volume_backup_createCmd.Flags().String("snapshot", "", "Snapshot to backup (name or ID)")
	volume_backup_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	volume_backupCmd.AddCommand(volume_backup_createCmd)
}
