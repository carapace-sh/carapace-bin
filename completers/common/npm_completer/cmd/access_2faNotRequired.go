package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/npm_completer/cmd/action"
	"github.com/spf13/cobra"
)

var access_2faNotRequiredCmd = &cobra.Command{
	Use:   "2fa-not-required",
	Short: "disable 2fa",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(access_2faNotRequiredCmd).Standalone()
	accessCmd.AddCommand(access_2faNotRequiredCmd)

	carapace.Gen(access_2faNotRequiredCmd).PositionalCompletion(
		action.ActionPackages(access_2faNotRequiredCmd),
	)
}
