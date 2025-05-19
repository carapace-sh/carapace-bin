package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tea"
	"github.com/spf13/cobra"
)

var logins_deleteCmd = &cobra.Command{
	Use:     "delete",
	Short:   "Remove a Gitea login",
	Aliases: []string{"rm"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logins_deleteCmd).Standalone()

	loginsCmd.AddCommand(logins_deleteCmd)

	carapace.Gen(logins_deleteCmd).PositionalCompletion(
		tea.ActionLogins(),
	)
}
