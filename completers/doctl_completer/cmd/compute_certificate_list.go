package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_certificate_listCmd = &cobra.Command{
	Use:   "list",
	Short: "Retrieve list of the account's stored certificates",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_certificate_listCmd).Standalone()
	compute_certificate_listCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `ID`, `Name`, `DNSNames`, `SHA1Fingerprint`, `NotAfter`, `Created`, `Type`, `State`")
	compute_certificate_listCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	compute_certificateCmd.AddCommand(compute_certificate_listCmd)

	carapace.Gen(compute_certificate_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("ID", "Name", "DNSNames", "SHA1Fingerprint", "NotAfter", "Created", "Type", "State").Invoke(c).Filter(c.Parts).ToA()
		}),
	})
}
