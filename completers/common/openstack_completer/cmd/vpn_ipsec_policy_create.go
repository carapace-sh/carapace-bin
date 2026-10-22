package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openstack_completer/cmd/action"
	"github.com/spf13/cobra"
)

var vpn_ipsec_policy_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create an IPsec policy",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpn_ipsec_policy_createCmd).Standalone()

	vpn_ipsec_policy_createCmd.Flags().String("auth-algorithm", "", "Authentication algorithm for IPsec policy")
	vpn_ipsec_policy_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	vpn_ipsec_policy_createCmd.Flags().String("description", "", "Description of the IPsec policy")
	vpn_ipsec_policy_createCmd.Flags().String("encapsulation-mode", "", "Encapsulation mode for IPsec policy")
	vpn_ipsec_policy_createCmd.Flags().String("encryption-algorithm", "", "Encryption algorithm for IPsec policy")
	vpn_ipsec_policy_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	vpn_ipsec_policy_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	vpn_ipsec_policy_createCmd.Flags().String("lifetime", "", "IPsec lifetime attributes.")
	vpn_ipsec_policy_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	vpn_ipsec_policy_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	vpn_ipsec_policy_createCmd.Flags().String("pfs", "", "Perfect Forward Secrecy for IPsec policy")
	vpn_ipsec_policy_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	vpn_ipsec_policy_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	vpn_ipsec_policy_createCmd.Flags().String("project", "", "Owner's project (name or ID)")
	vpn_ipsec_policy_createCmd.Flags().String("project-domain", "", "Domain the project belongs to (name or ID).")
	vpn_ipsec_policy_createCmd.Flags().String("transform-protocol", "", "Transform protocol for IPsec policy")
	vpn_ipsec_policy_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	vpn_ipsec_policyCmd.AddCommand(vpn_ipsec_policy_createCmd)

	carapace.Gen(vpn_ipsec_policy_createCmd).FlagCompletion(carapace.ActionMap{
		"auth-algorithm":       carapace.ActionValues("sha1", "sha256", "sha384", "sha512", "aes-xcbc", "aes-cmac"),
		"column":               action.ActionColumns(vpn_ipsec_policy_createCmd),
		"encapsulation-mode":   carapace.ActionValues("tunnel", "transport"),
		"encryption-algorithm": carapace.ActionValues("3des", "aes-128", "aes-192", "aes-256", "aes-128-ccm-8", "aes-192-ccm-8", "aes-256-ccm-8", "aes-128-ccm-12", "aes-192-ccm-12", "aes-256-ccm-12", "aes-128-ccm-16", "aes-192-ccm-16", "aes-256-ccm-16", "aes-128-gcm-8", "aes-192-gcm-8", "aes-256-gcm-8", "aes-128-gcm-12", "aes-192-gcm-12", "aes-256-gcm-12", "aes-128-gcm-16", "aes-192-gcm-16", "aes-256-gcm-16", "aes-128-ctr", "aes-192-ctr", "aes-256-ctr"),
		"format":               carapace.ActionValues("json", "shell", "table", "value", "yaml"),
		"pfs":                  carapace.ActionValues("group2", "group5", "group14", "group15", "group16", "group17", "group18", "group19", "group20", "group21", "group22", "group23", "group24", "group25", "group26", "group27", "group28", "group29", "group30", "group31"),
		"project":              action.ActionProjects(),
		"project-domain":       action.ActionDomains(),
		"transform-protocol":   carapace.ActionValues("esp", "ah", "ah-esp"),
	})
}
