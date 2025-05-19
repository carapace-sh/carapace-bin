package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/systemctl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var serviceLogTargetCmd = &cobra.Command{
	Use:     "service-log-target",
	Short:   "Get/set logging target for service",
	GroupID: "unit",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serviceLogTargetCmd).Standalone()

	rootCmd.AddCommand(serviceLogTargetCmd)

	carapace.Gen(serviceLogTargetCmd).PositionalCompletion(
		action.ActionServices(serviceLogTargetCmd),
		action.ActionTargets(serviceLogTargetCmd),
	)
}
