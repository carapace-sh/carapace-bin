package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_cdn_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List CDNs that have already been created",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_cdn_listCmd).Standalone()
	compute_cdn_listCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `ID`, `Origin`, `Endpoint`, `TTL`, `CustomDomain`, `CertificateID`, `CreatedAt`")
	compute_cdn_listCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	compute_cdnCmd.AddCommand(compute_cdn_listCmd)

	carapace.Gen(compute_cdn_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("ID", "Origin", "Endpoint", "TTL", "CustomDomain", "CertificateID", "CreatedAt").Invoke(c).Filter(c.Parts).ToA()
		}),
	})
}
