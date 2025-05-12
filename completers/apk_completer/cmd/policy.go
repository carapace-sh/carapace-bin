package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/apk_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/completers/apk_completer/cmd/common"
	"github.com/spf13/cobra"
)

var policyCmd = &cobra.Command{
	Use:     "policy",
	Short:   "Show repository policy for packages",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: "querying package information",
}

func init() {
	carapace.Gen(policyCmd).Standalone()

	common.AddGlobalFlags(policyCmd)
	rootCmd.AddCommand(policyCmd)

	carapace.Gen(policyCmd).PositionalAnyCompletion(
		action.ActionPackageSearch(policyCmd).FilterArgs(),
	)
}
