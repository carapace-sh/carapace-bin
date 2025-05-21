package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/fnm_completer/cmd/action"
	"github.com/spf13/cobra"
)

var aliasCmd = &cobra.Command{
	Use:   "alias",
	Short: "Alias a version to a common name",
	Run:   func(*cobra.Command, []string) {},
}

func init() {
	rootCmd.AddCommand(aliasCmd)

	addCommonFlags(aliasCmd)

	carapace.Gen(aliasCmd).Standalone()

	carapace.Gen(aliasCmd).PositionalCompletion(
		action.ActionLocalVersions(),
	)
}
