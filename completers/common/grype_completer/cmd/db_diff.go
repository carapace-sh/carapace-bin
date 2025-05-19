package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var db_diffCmd = &cobra.Command{
	Use:   "diff [flags] base_db_url target_db_url",
	Short: "diff two DBs and display the result",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(db_diffCmd).Standalone()

	db_diffCmd.Flags().BoolP("delete", "d", false, "delete downloaded databases after diff occurs")
	db_diffCmd.Flags().StringP("output", "o", "table", "format to display results")
	dbCmd.AddCommand(db_diffCmd)

	carapace.Gen(db_diffCmd).FlagCompletion(carapace.ActionMap{
		"output": carapace.ActionValues("table", "json"),
	})
}
