package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var databases_pool_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List connection pools for a database cluster",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(databases_pool_listCmd).Standalone()
	databases_pool_listCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `User`, `Name`, `Size`, `Database`, `Mode`, `URI`")
	databases_pool_listCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	databases_poolCmd.AddCommand(databases_pool_listCmd)

	carapace.Gen(databases_pool_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("User", "Name", "Size", "Database", "Mode", "URI").Invoke(c).Filter(c.Parts).ToA()
		}),
	})
}
