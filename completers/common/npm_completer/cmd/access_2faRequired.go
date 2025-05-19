package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/npm_completer/cmd/action"
	"github.com/spf13/cobra"
)

var access_2faRequiredCmd = &cobra.Command{
	Use:   "2fa-required",
	Short: "enable 2fa",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(access_2faRequiredCmd).Standalone()
	accessCmd.AddCommand(access_2faRequiredCmd)

	carapace.Gen(access_2faRequiredCmd).PositionalCompletion(
		action.ActionPackages(access_2faRequiredCmd),
	)
}
