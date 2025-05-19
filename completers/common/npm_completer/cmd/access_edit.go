package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/npm_completer/cmd/action"
	"github.com/spf13/cobra"
)

var access_editCmd = &cobra.Command{
	Use:   "edit",
	Short: "edit access",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(access_editCmd).Standalone()
	accessCmd.AddCommand(access_editCmd)

	carapace.Gen(access_editCmd).PositionalCompletion(
		action.ActionPackages(access_editCmd),
	)
}
