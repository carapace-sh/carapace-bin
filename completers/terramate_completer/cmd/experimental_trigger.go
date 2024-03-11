package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/terramate"
	"github.com/spf13/cobra"
)

var experimental_triggerCmd = &cobra.Command{
	Use:   "trigger",
	Short: "Triggers a stack",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(experimental_triggerCmd).Standalone()

	experimental_triggerCmd.Flags().String("cloud-status", "", "Filter by status")
	experimental_triggerCmd.Flags().String("experimental-status", "", "Filter by status (Deprecated)")
	experimental_triggerCmd.Flags().String("reason", "", "Reason for the stack being triggered")
	experimentalCmd.AddCommand(experimental_triggerCmd)

	carapace.Gen(experimental_triggerCmd).FlagCompletion(carapace.ActionMap{
		"cloud-status": terramate.ActionCloudStatus(),
	})

	carapace.Gen(experimental_triggerCmd).PositionalCompletion(
		terramate.ActionStacks(),
	)
}
