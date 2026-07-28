package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var downCmd = &cobra.Command{
	Use:   "down",
	Short: "Disconnect from Tailscale",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(downCmd).Standalone()

	downCmd.Flags().String("accept-risk", "", "accept risk and skip confirmation for risk types")
	downCmd.Flags().String("reason", "", "reason for the disconnect, if required by a policy")
	rootCmd.AddCommand(downCmd)

	carapace.Gen(downCmd).FlagCompletion(carapace.ActionMap{
		"accept-risk": carapace.ActionValues("lose-ssh", "mac-app-connector", "all"),
	})
}
