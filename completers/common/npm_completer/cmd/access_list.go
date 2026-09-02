package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/npm_completer/cmd/action"
	"github.com/spf13/cobra"
)

var access_listCmd = &cobra.Command{
	Use:   "list",
	Short: "list packages or collaborators",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(access_listCmd).Standalone()
	accessCmd.AddCommand(access_listCmd)

	carapace.Gen(access_listCmd).PositionalCompletion(
		carapace.ActionValues("packages", "collaborators"),
	)
	carapace.Gen(access_listCmd).PositionalAnyCompletion(
		action.ActionPackages(access_listCmd),
	)
}
