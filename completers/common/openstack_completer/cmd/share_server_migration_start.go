package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_server_migration_startCmd = &cobra.Command{
	Use:   "start",
	Short: "Migrates share server to a new host.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_server_migration_startCmd).Standalone()

	share_server_migration_startCmd.Flags().Bool("check-only", false, "Run a dry-run of the share server migration. ")
	share_server_migration_startCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	share_server_migration_startCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	share_server_migration_startCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	share_server_migration_startCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	share_server_migration_startCmd.Flags().String("new-share-network", "", "Specify a new share network for the share server.")
	share_server_migration_startCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	share_server_migration_startCmd.Flags().String("nondisruptive", "", "Enforces migration to be nondisruptive.")
	share_server_migration_startCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	share_server_migration_startCmd.Flags().String("preserve-snapshots", "", "Set to True if snapshots must be preserved at the migration destination.")
	share_server_migration_startCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	share_server_migration_startCmd.Flags().String("variable", "", "==SUPPRESS==")
	share_server_migration_startCmd.Flags().String("writable", "", "Enforces migration to keep all its shares writable while contents are being moved.")
	share_server_migration_startCmd.MarkFlagRequired("nondisruptive")
	share_server_migration_startCmd.MarkFlagRequired("preserve-snapshots")
	share_server_migration_startCmd.MarkFlagRequired("writable")
	share_server_migrationCmd.AddCommand(share_server_migration_startCmd)
}
