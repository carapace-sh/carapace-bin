package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/npm_completer/cmd/action"
	"github.com/spf13/cobra"
)

var access_getCmd = &cobra.Command{
	Use:   "get",
	Short: "get access status",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(access_getCmd).Standalone()
	accessCmd.AddCommand(access_getCmd)

	carapace.Gen(access_getCmd).PositionalCompletion(
		carapace.ActionValues("status"),
	)
	carapace.Gen(access_getCmd).PositionalAnyCompletion(
		action.ActionPackages(access_getCmd),
	)
}
