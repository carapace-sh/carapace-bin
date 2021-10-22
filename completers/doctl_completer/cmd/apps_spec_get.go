package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var apps_spec_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve an application's spec",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apps_spec_getCmd).Standalone()
	apps_spec_getCmd.Flags().String("deployment", "", "optional: a deployment ID")
	apps_spec_getCmd.Flags().String("format", "yaml", "the format to output the spec in; either \"yaml\" or \"json\"")
	apps_specCmd.AddCommand(apps_spec_getCmd)

	// TODO deployment
	carapace.Gen(apps_spec_getCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("yaml", "json"),
	})

	// TODO positional
}
