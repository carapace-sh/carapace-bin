package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_certificate_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve details about a certificate",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_certificate_getCmd).Standalone()
	compute_certificate_getCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `ID`, `Name`, `DNSNames`, `SHA1Fingerprint`, `NotAfter`, `Created`, `Type`, `State`")
	compute_certificate_getCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	compute_certificateCmd.AddCommand(compute_certificate_getCmd)
}
