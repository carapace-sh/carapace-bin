package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/helm_completer/cmd/action"
	"github.com/spf13/cobra"
)

var get_allCmd = &cobra.Command{
	Use:   "all",
	Short: "download all information for a named release",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(get_allCmd).Standalone()
	get_allCmd.Flags().Int("revision", 0, "get the named release with revision")
	get_allCmd.Flags().String("template", "", "go template for formatting the output, eg: {{.Release.Name}}")
	getCmd.AddCommand(get_allCmd)

	carapace.Gen(get_allCmd).PositionalCompletion(
		action.ActionReleases(get_allCmd),
	)
}
