package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var databases_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List your database clusters",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(databases_listCmd).Standalone()
	databases_listCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `ID`, `Name`, `Engine`, `Version`, `NumNodes`, `Region`, `Status`, `Size`, `URI`, `Created`")
	databases_listCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	databasesCmd.AddCommand(databases_listCmd)

	carapace.Gen(databases_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("ID", "Name", "Engine", "Version", "NumNodes", "Region", "Status", "Size", "URI", "Created").Invoke(c).Filter(c.Parts).ToA()
		}),
	})
}
