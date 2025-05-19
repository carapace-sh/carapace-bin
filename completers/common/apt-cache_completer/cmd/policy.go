package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/apt-cache_completer/cmd/action"
	"github.com/spf13/cobra"
)

var policyCmd = &cobra.Command{
	Use:   "policy",
	Short: "debug issues relating to the preferences file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(policyCmd).Standalone()

	rootCmd.AddCommand(policyCmd)

	carapace.Gen(policyCmd).PositionalAnyCompletion(
		action.ActionPackages(),
	)
}
