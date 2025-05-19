package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/helm_completer/cmd/action"
	"github.com/spf13/cobra"
)

var get_manifestCmd = &cobra.Command{
	Use:   "manifest",
	Short: "download the manifest for a named release",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(get_manifestCmd).Standalone()
	get_manifestCmd.Flags().Int("revision", 0, "get the named release with revision")
	getCmd.AddCommand(get_manifestCmd)

	carapace.Gen(get_manifestCmd).PositionalCompletion(
		action.ActionReleases(get_manifestCmd),
	)
}
