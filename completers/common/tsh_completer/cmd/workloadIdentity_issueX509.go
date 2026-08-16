package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workloadIdentity_issueX509Cmd = &cobra.Command{
	Use:   "issue-x509",
	Short: "Use Teleport Workload Identity to issue an X509 credential and write it to a local directory.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workloadIdentity_issueX509Cmd).Standalone()

	workloadIdentity_issueX509Cmd.Flags().String("credential-ttl", "1h", "Sets the time to live for the credential.")
	workloadIdentity_issueX509Cmd.Flags().String("label-selector", "", "A label-based selector for which workload identities to issue. Multiple labels can be provided using ','.")
	workloadIdentity_issueX509Cmd.Flags().String("name-selector", "", "The name of the workload identity to issue.")
	workloadIdentity_issueX509Cmd.Flags().String("output", "", "Path to the directory to write the SVID into.")
	workloadIdentity_issueX509Cmd.MarkFlagRequired("output")
	workloadIdentityCmd.AddCommand(workloadIdentity_issueX509Cmd)
}
