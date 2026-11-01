package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var agentCmd = &cobra.Command{
	Use:   "agent",
	Short: "NetworkManager secret agent or polkit agent",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(agentCmd).Standalone()

	rootCmd.AddCommand(agentCmd)

	carapace.Gen(agentCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"all", "Runs nmcli as both NetworkManager secret and a polkit agent.",
			"polkit", "Registers nmcli as a polkit action for the user session",
			"secret", "Runs nmcli as NetworkManager secret agent",
		),
	)
}
