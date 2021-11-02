package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var databases_backupsCmd = &cobra.Command{
	Use:   "backups",
	Short: "List database cluster backups",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(databases_backupsCmd).Standalone()
	databases_backupsCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `Size`, `Created`")
	databases_backupsCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	databasesCmd.AddCommand(databases_backupsCmd)

	carapace.Gen(databases_backupsCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("Size", "Created").Invoke(c).Filter(c.Parts).ToA()
		}),
	})
}
