package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var server_backup_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a server backup image",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(server_backup_createCmd).Standalone()

	server_backup_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	server_backup_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	server_backup_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	server_backup_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	server_backup_createCmd.Flags().String("name", "", "Name of the backup image (default: server name)")
	server_backup_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	server_backup_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	server_backup_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	server_backup_createCmd.Flags().String("rotate", "", "Number of backups to keep (default: 1)")
	server_backup_createCmd.Flags().String("type", "", "Used to populate the backup_type property of the backup image (default: empty)")
	server_backup_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	server_backup_createCmd.Flags().Bool("wait", false, "Wait for backup image create to complete")
	server_backupCmd.AddCommand(server_backup_createCmd)
}
