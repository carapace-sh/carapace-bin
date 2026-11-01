package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/bluetoothctl"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var agentCmd = &cobra.Command{
	Use:   "agent <on|off|auto|capability>",
	Short: "Enable/disable agent with given capability",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(agentCmd).Standalone()
	rootCmd.AddCommand(agentCmd)
	carapace.Gen(agentCmd).PositionalCompletion(
		carapace.Batch(
			carapace.ActionValues("auto", "off", "on").StyleF(style.ForKeyword),
			bluetoothctl.ActionAgentCapabilities(),
		).ToA(),
	)
}
