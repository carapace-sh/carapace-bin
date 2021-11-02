package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_cdn_updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update the configuration for a CDN",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_cdn_updateCmd).Standalone()
	compute_cdn_updateCmd.Flags().String("certificate-id", "", "Specify a Certificate ID for the custom domain")
	compute_cdn_updateCmd.Flags().String("domain", "", "Specify a custom domain to use with the CDN")
	compute_cdn_updateCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `ID`, `Origin`, `Endpoint`, `TTL`, `CustomDomain`, `CertificateID`, `CreatedAt`")
	compute_cdn_updateCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	compute_cdn_updateCmd.Flags().Int("ttl", 3600, "The \"Time To Live\" (TTL) value for cached content, in seconds")
	compute_cdnCmd.AddCommand(compute_cdn_updateCmd)

	carapace.Gen(compute_cdn_updateCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues(`ID`, `Origin`, `Endpoint`, `TTL`, `CustomDomain`, `CertificateID`, `CreatedAt`).Invoke(c).Filter(c.Parts).ToA()
		}),
	})
}
