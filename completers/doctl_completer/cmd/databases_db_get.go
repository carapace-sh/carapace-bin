package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var databases_db_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve the name of a database within a cluster",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(databases_db_getCmd).Standalone()
	databases_db_getCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `Name`")
	databases_db_getCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	databases_dbCmd.AddCommand(databases_db_getCmd)

	carapace.Gen(databases_db_getCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("Name").Invoke(c).Filter(c.Parts).ToA()
		}),
	})
}
