package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workloadIdentity_issueJwtCmd = &cobra.Command{
	Use:   "issue-jwt",
	Short: "Use Teleport Workload Identity to issue a JWT credential and write it to a local directory.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workloadIdentity_issueJwtCmd).Standalone()

	workloadIdentity_issueJwtCmd.Flags().String("audience", "", "The audience to include in the JWT. Can be specified multiple times.")
	workloadIdentity_issueJwtCmd.Flags().String("credential-ttl", "1h", "Sets the time to live for the credential.")
	workloadIdentity_issueJwtCmd.Flags().String("label-selector", "", "A label-based selector for which workload identities to issue. Multiple labels can be provided using ','.")
	workloadIdentity_issueJwtCmd.Flags().String("name-selector", "", "The name of the workload identity to issue.")
	workloadIdentity_issueJwtCmd.Flags().String("output", "", "Path to the directory to write the SVID into.")
	workloadIdentity_issueJwtCmd.MarkFlagRequired("audience")
	workloadIdentity_issueJwtCmd.MarkFlagRequired("output")
	workloadIdentityCmd.AddCommand(workloadIdentity_issueJwtCmd)
}
