package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_cdn_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a CDN",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_cdn_createCmd).Standalone()
	compute_cdn_createCmd.Flags().String("certificate-id", "", "Specify a Certificate ID for the custom domain")
	compute_cdn_createCmd.Flags().String("domain", "", "Specify a custom domain to use with the CDN")
	compute_cdn_createCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `ID`, `Origin`, `Endpoint`, `TTL`, `CustomDomain`, `CertificateID`, `CreatedAt`")
	compute_cdn_createCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	compute_cdn_createCmd.Flags().Int("ttl", 3600, "The \"Time To Live\" (TTL) value for cached content, in seconds")
	compute_cdnCmd.AddCommand(compute_cdn_createCmd)

	carapace.Gen(compute_cdn_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("ID", "Origin", "Endpoint", "TTL", "CustomDomain", "CertificateID", "CreatedAt").Invoke(c).Filter(c.Parts).ToA()
		}),
	})
}
