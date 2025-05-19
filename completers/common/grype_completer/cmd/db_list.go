package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var db_listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all DBs available according to the listing URL",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(db_listCmd).Standalone()

	db_listCmd.Flags().StringP("output", "o", "text", "format to display results")
	dbCmd.AddCommand(db_listCmd)

	carapace.Gen(db_listCmd).FlagCompletion(carapace.ActionMap{
		"output": carapace.ActionValues("text", "raw", "json"),
	})
}
