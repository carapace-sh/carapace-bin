package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workloadIdentityCmd = &cobra.Command{
	Use:   "workload-identity",
	Short: "Issue Workload Identity credentials.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workloadIdentityCmd).Standalone()

	rootCmd.AddCommand(workloadIdentityCmd)
}
