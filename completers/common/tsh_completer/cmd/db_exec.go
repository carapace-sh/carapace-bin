package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var db_execCmd = &cobra.Command{
	Use:   "exec",
	Short: "Execute database commands on target database services.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(db_execCmd).Standalone()

	db_execCmd.Flags().Bool("confirm", true, "Confirm selected database services before executing command.")
	db_execCmd.Flags().StringP("db-name", "n", "", "Database name to log in to.")
	db_execCmd.Flags().StringP("db-roles", "r", "", "List of comma separate database roles to use for auto-provisioned user.")
	db_execCmd.Flags().StringP("db-user", "u", "", "Database user to log in as.")
	db_execCmd.Flags().String("dbs", "", "List of comma separated target database services. Mutually exclusive with --search or --labels.")
	db_execCmd.Flags().String("labels", "", "List of comma separated labels to filter by labels (e.g. key1=value1,key2=value2).")
	db_execCmd.Flags().Bool("no-confirm", false, "Confirm selected database services before executing command.")
	db_execCmd.Flags().String("output-dir", "", "Directory to store command output per target database service. A summary is saved as \"summary.json\".")
	db_execCmd.Flags().String("parallel", "1", "Run commands on target databases in parallel. Defaults to 1, and maximum allowed is 10.")
	db_execCmd.Flags().String("search", "", "List of comma separated search keywords or phrases enclosed in quotations (e.g. --search=foo,bar,\"some phrase\").")
	db_execCmd.Flag("no-confirm").Hidden = true
	dbCmd.AddCommand(db_execCmd)
}
