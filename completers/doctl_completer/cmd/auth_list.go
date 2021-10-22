package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var auth_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available authentication contexts",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auth_listCmd).Standalone()
	auth_listCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `text`")
	authCmd.AddCommand(auth_listCmd)

	carapace.Gen(apps_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("text").Invoke(c).Filter(c.Parts).ToA()
		}),
	})
}
