package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var svid_issueCmd = &cobra.Command{
	Use:   "issue",
	Short: "Issue a SPIFFE SVID using Teleport Workload Identity and write it to a local directory.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(svid_issueCmd).Standalone()

	svid_issueCmd.Flags().String("dns-san", "", "DNS SANs to include in the SVID. By default, none are included.")
	svid_issueCmd.Flags().String("ip-san", "", "IP SANs to include in the SVID. By default, none are included.")
	svid_issueCmd.Flags().String("output", "", "Path to the directory to write the SVID into.")
	svid_issueCmd.Flags().String("svid-ttl", "1h", "Sets the time to live for the SVID.")
	svid_issueCmd.Flags().String("type", "x509", "Type of the SVID to issue (x509). Defaults to x509.")
	svid_issueCmd.MarkFlagRequired("output")
	svidCmd.AddCommand(svid_issueCmd)
}
