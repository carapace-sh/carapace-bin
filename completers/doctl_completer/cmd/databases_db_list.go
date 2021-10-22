package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var databases_db_listCmd = &cobra.Command{
	Use:   "list",
	Short: "Retrieve a list of databases within a cluster",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(databases_db_listCmd).Standalone()
	databases_db_listCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `Name`")
	databases_db_listCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	databases_dbCmd.AddCommand(databases_db_listCmd)

	carapace.Gen(databases_db_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("Name").Invoke(c).Filter(c.Parts).ToA()
		}),
	})
}
