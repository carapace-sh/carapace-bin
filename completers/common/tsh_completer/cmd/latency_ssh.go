package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var latency_sshCmd = &cobra.Command{
	Use:   "ssh",
	Short: "Measure latency to a particular SSH host.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(latency_sshCmd).Standalone()

	latency_sshCmd.Flags().StringP("cluster", "c", "", "Specify the Teleport cluster to connect.")
	latency_sshCmd.Flags().Bool("no-no-resume", false, "Disable SSH connection resumption.")
	latency_sshCmd.Flags().Bool("no-resume", false, "Disable SSH connection resumption.")
	latency_sshCmd.Flag("no-no-resume").Hidden = true
	latencyCmd.AddCommand(latency_sshCmd)
}
