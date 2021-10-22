package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var databases_pool_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve information about a database connection pool",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(databases_pool_getCmd).Standalone()
	databases_pool_getCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `User`, `Name`, `Size`, `Database`, `Mode`, `URI`")
	databases_pool_getCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	databases_poolCmd.AddCommand(databases_pool_getCmd)

	carapace.Gen(databases_pool_getCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("User", "Name", "Size", "Database", "Mode", "URI").Invoke(c).Filter(c.Parts).ToA()
		}),
	})
}
