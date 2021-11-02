package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var databases_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get details for a database cluster",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(databases_getCmd).Standalone()
	databases_getCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `ID`, `Name`, `Engine`, `Version`, `NumNodes`, `Region`, `Status`, `Size`, `URI`, `Created`")
	databases_getCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	databasesCmd.AddCommand(databases_getCmd)

	carapace.Gen(databases_getCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("ID", "Name", "Engine", "Version", "NumNodes", "Region", "Status", "Size", "URI", "Created").Invoke(c).Filter(c.Parts).ToA()
		}),
	})
}
