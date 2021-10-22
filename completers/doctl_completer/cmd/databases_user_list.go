package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var databases_user_listCmd = &cobra.Command{
	Use:   "list",
	Short: "Retrieve list of database users",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(databases_user_listCmd).Standalone()
	databases_user_listCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `Name`, `Role`, `Password`")
	databases_user_listCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	databases_userCmd.AddCommand(databases_user_listCmd)
}
