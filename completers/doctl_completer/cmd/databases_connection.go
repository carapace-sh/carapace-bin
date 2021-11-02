package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var databases_connectionCmd = &cobra.Command{
	Use:   "connection",
	Short: "Retrieve connection details for a database cluster",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(databases_connectionCmd).Standalone()
	databases_connectionCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `URI`, `Database`, `Host`, `Port`, `User`, `Password`, `SSL`")
	databases_connectionCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	databasesCmd.AddCommand(databases_connectionCmd)

	carapace.Gen(databases_connectionCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("URI", "Database", "Host", "Port", "User", "Password", "SSL").Invoke(c).Filter(c.Parts).ToA()
		}),
	})
}
