package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bridge_transformations_listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List available transformations",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bridge_transformations_listCmd).Standalone()

	bridge_transformations_listCmd.Flags().String("format", "", "Format the output. Values: [table | json]")
	bridge_transformations_listCmd.Flags().BoolP("quiet", "q", false, "Only display transformer names")
	bridge_transformationsCmd.AddCommand(bridge_transformations_listCmd)

	carapace.Gen(bridge_transformations_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("table", "json"),
	})
}
