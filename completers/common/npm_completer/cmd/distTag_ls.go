package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/npm_completer/cmd/action"
	"github.com/spf13/cobra"
)

var distTag_lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "show all the dist-tags for a package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(distTag_lsCmd).Standalone()

	distTagCmd.AddCommand(distTag_lsCmd)

	carapace.Gen(distTag_lsCmd).PositionalCompletion(
		action.ActionPackageNames(distTag_lsCmd),
	)
}
