package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/security"
	"github.com/spf13/cobra"
)

var addCertificatesCmd = &cobra.Command{
	Use:   "add-certificates",
	Short: "Add certificates to a keychain",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var addGenericPasswordCmd = &cobra.Command{
	Use:   "add-generic-password",
	Short: "Add a generic password item",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var addInternetPasswordCmd = &cobra.Command{
	Use:   "add-internet-password",
	Short: "Add an internet password item",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var deleteGenericPasswordCmd = &cobra.Command{
	Use:   "delete-generic-password",
	Short: "Delete a generic password item",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var deleteInternetPasswordCmd = &cobra.Command{
	Use:   "delete-internet-password",
	Short: "Delete an internet password item",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var deleteCertificateCmd = &cobra.Command{
	Use:   "delete-certificate",
	Short: "Delete a certificate from a keychain",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var findCertificateCmd = &cobra.Command{
	Use:   "find-certificate",
	Short: "Find a certificate item",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var findGenericPasswordCmd = &cobra.Command{
	Use:   "find-generic-password",
	Short: "Find a generic password item",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var findInternetPasswordCmd = &cobra.Command{
	Use:   "find-internet-password",
	Short: "Find an internet password item",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var findIdentityCmd = &cobra.Command{
	Use:   "find-identity",
	Short: "Find an identity (certificate + private key)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(addCertificatesCmd).Standalone()
	addCertificatesCmd.Flags().StringP("keychain", "k", "", "Use keychain rather than the default keychain")
	rootCmd.AddCommand(addCertificatesCmd)
	carapace.Gen(addCertificatesCmd).PositionalCompletion(carapace.ActionFiles())
	carapace.Gen(addCertificatesCmd).FlagCompletion(carapace.ActionMap{
		"keychain": security.ActionKeychains(),
	})

	carapace.Gen(addGenericPasswordCmd).Standalone()
	addGenericPasswordCmd.Flags().StringP("account", "a", "", "Specify account name (required)")
	addGenericPasswordCmd.Flags().BoolP("allow-all", "A", false, "Allow any application to access without warning")
	addGenericPasswordCmd.Flags().StringP("attribute", "G", "", "Specify generic attribute value")
	addGenericPasswordCmd.Flags().StringP("comment", "j", "", "Specify comment string")
	addGenericPasswordCmd.Flags().StringP("creator", "c", "", "Specify item creator (four-character code)")
	addGenericPasswordCmd.Flags().StringP("kind", "D", "", "Specify kind (default: application password)")
	addGenericPasswordCmd.Flags().StringP("label", "l", "", "Specify label (default: service name)")
	addGenericPasswordCmd.Flags().StringP("password", "w", "", "Specify password to be added")
	addGenericPasswordCmd.Flags().StringP("service", "s", "", "Specify service name (required)")
	addGenericPasswordCmd.Flags().StringP("tool", "T", "", "Specify an application which may access this item")
	addGenericPasswordCmd.Flags().StringP("type", "C", "", "Specify item type (four-character code)")
	addGenericPasswordCmd.Flags().BoolP("update", "U", false, "Update item if it already exists")
	rootCmd.AddCommand(addGenericPasswordCmd)

	carapace.Gen(addInternetPasswordCmd).Standalone()
	addInternetPasswordCmd.Flags().StringP("account", "a", "", "Specify account name (required)")
	addInternetPasswordCmd.Flags().BoolP("allow-all", "A", false, "Allow any application to access without warning")
	addInternetPasswordCmd.Flags().StringP("auth-type", "t", "", "Specify authentication type (default: dflt)")
	addInternetPasswordCmd.Flags().StringP("comment", "j", "", "Specify comment string")
	addInternetPasswordCmd.Flags().StringP("creator", "c", "", "Specify item creator (four-character code)")
	addInternetPasswordCmd.Flags().StringP("domain", "d", "", "Specify security domain string")
	addInternetPasswordCmd.Flags().StringP("kind", "D", "", "Specify kind (default: application password)")
	addInternetPasswordCmd.Flags().StringP("label", "l", "", "Specify label (default: service name)")
	addInternetPasswordCmd.Flags().StringP("password", "w", "", "Specify password to be added")
	addInternetPasswordCmd.Flags().StringP("path", "p", "", "Specify path string")
	addInternetPasswordCmd.Flags().StringP("port", "P", "", "Specify port number")
	addInternetPasswordCmd.Flags().StringP("protocol", "r", "", "Specify protocol (four-character code)")
	addInternetPasswordCmd.Flags().StringP("server", "s", "", "Specify server name (required)")
	addInternetPasswordCmd.Flags().StringP("tool", "T", "", "Specify an application which may access this item")
	addInternetPasswordCmd.Flags().StringP("type", "C", "", "Specify item type (four-character code)")
	addInternetPasswordCmd.Flags().BoolP("update", "U", false, "Update item if it already exists")
	rootCmd.AddCommand(addInternetPasswordCmd)

	carapace.Gen(deleteGenericPasswordCmd).Standalone()
	deleteGenericPasswordCmd.Flags().StringP("account", "a", "", "Match account string")
	deleteGenericPasswordCmd.Flags().StringP("attribute", "G", "", "Match value string")
	deleteGenericPasswordCmd.Flags().StringP("comment", "j", "", "Match comment string")
	deleteGenericPasswordCmd.Flags().StringP("creator", "c", "", "Match creator (four-character code)")
	deleteGenericPasswordCmd.Flags().StringP("kind", "D", "", "Match kind string")
	deleteGenericPasswordCmd.Flags().StringP("label", "l", "", "Match label string")
	deleteGenericPasswordCmd.Flags().StringP("service", "s", "", "Match service string")
	deleteGenericPasswordCmd.Flags().StringP("type", "C", "", "Match type (four-character code)")
	rootCmd.AddCommand(deleteGenericPasswordCmd)

	carapace.Gen(deleteInternetPasswordCmd).Standalone()
	deleteInternetPasswordCmd.Flags().StringP("account", "a", "", "Match account string")
	deleteInternetPasswordCmd.Flags().StringP("auth-type", "t", "", "Match authentication type")
	deleteInternetPasswordCmd.Flags().StringP("comment", "j", "", "Match comment string")
	deleteInternetPasswordCmd.Flags().StringP("creator", "c", "", "Match creator (four-character code)")
	deleteInternetPasswordCmd.Flags().StringP("domain", "d", "", "Match security domain string")
	deleteInternetPasswordCmd.Flags().StringP("kind", "D", "", "Match kind string")
	deleteInternetPasswordCmd.Flags().StringP("label", "l", "", "Match label string")
	deleteInternetPasswordCmd.Flags().StringP("path", "p", "", "Match path string")
	deleteInternetPasswordCmd.Flags().StringP("port", "P", "", "Match port number")
	deleteInternetPasswordCmd.Flags().StringP("protocol", "r", "", "Match protocol (four-character code)")
	deleteInternetPasswordCmd.Flags().StringP("server", "s", "", "Match server string")
	deleteInternetPasswordCmd.Flags().StringP("type", "C", "", "Match type (four-character code)")
	rootCmd.AddCommand(deleteInternetPasswordCmd)

	carapace.Gen(deleteCertificateCmd).Standalone()
	deleteCertificateCmd.Flags().StringP("hash", "Z", "", "Specify certificate to delete by SHA-1 hash")
	deleteCertificateCmd.Flags().StringP("name", "c", "", "Specify certificate to delete by common name")
	deleteCertificateCmd.Flags().BoolP("trust", "t", false, "Also delete user trust settings for this certificate")
	rootCmd.AddCommand(deleteCertificateCmd)

	carapace.Gen(findCertificateCmd).Standalone()
	findCertificateCmd.Flags().BoolP("all", "a", false, "Find all matching certificates, not just the first")
	findCertificateCmd.Flags().StringP("email", "e", "", "Match on email address when searching")
	findCertificateCmd.Flags().BoolP("email-addresses", "m", false, "Show the email addresses in the certificate")
	findCertificateCmd.Flags().BoolP("hash", "Z", false, "Print SHA-1 hash of the certificate")
	findCertificateCmd.Flags().StringP("name", "c", "", "Match on name when searching")
	findCertificateCmd.Flags().BoolP("pem", "p", false, "Output certificate in PEM format")
	rootCmd.AddCommand(findCertificateCmd)

	carapace.Gen(findGenericPasswordCmd).Standalone()
	findGenericPasswordCmd.Flags().StringP("account", "a", "", "Match account string")
	findGenericPasswordCmd.Flags().StringP("attribute", "G", "", "Match value string")
	findGenericPasswordCmd.Flags().StringP("comment", "j", "", "Match comment string")
	findGenericPasswordCmd.Flags().StringP("creator", "c", "", "Match creator (four-character code)")
	findGenericPasswordCmd.Flags().BoolP("get-password", "g", false, "Display the password for the item found")
	findGenericPasswordCmd.Flags().StringP("kind", "D", "", "Match kind string")
	findGenericPasswordCmd.Flags().StringP("label", "l", "", "Match label string")
	findGenericPasswordCmd.Flags().BoolP("password-only", "w", false, "Display the password only for the item found")
	findGenericPasswordCmd.Flags().StringP("service", "s", "", "Match service string")
	findGenericPasswordCmd.Flags().StringP("type", "C", "", "Match type (four-character code)")
	rootCmd.AddCommand(findGenericPasswordCmd)

	carapace.Gen(findInternetPasswordCmd).Standalone()
	findInternetPasswordCmd.Flags().StringP("account", "a", "", "Match account string")
	findInternetPasswordCmd.Flags().StringP("auth-type", "t", "", "Match authentication type")
	findInternetPasswordCmd.Flags().StringP("comment", "j", "", "Match comment string")
	findInternetPasswordCmd.Flags().StringP("creator", "c", "", "Match creator (four-character code)")
	findInternetPasswordCmd.Flags().StringP("domain", "d", "", "Match security domain string")
	findInternetPasswordCmd.Flags().BoolP("get-password", "g", false, "Display the password for the item found")
	findInternetPasswordCmd.Flags().StringP("kind", "D", "", "Match kind string")
	findInternetPasswordCmd.Flags().StringP("label", "l", "", "Match label string")
	findInternetPasswordCmd.Flags().BoolP("password-only", "w", false, "Display the password only for the item found")
	findInternetPasswordCmd.Flags().StringP("path", "p", "", "Match path string")
	findInternetPasswordCmd.Flags().StringP("port", "P", "", "Match port number")
	findInternetPasswordCmd.Flags().StringP("protocol", "r", "", "Match protocol (four-character code)")
	findInternetPasswordCmd.Flags().StringP("server", "s", "", "Match server string")
	findInternetPasswordCmd.Flags().StringP("type", "C", "", "Match type (four-character code)")
	rootCmd.AddCommand(findInternetPasswordCmd)

	carapace.Gen(findIdentityCmd).Standalone()
	findIdentityCmd.Flags().StringP("policy", "p", "", "Specify policy (basic, ssl-client, ssl-server, smime, etc.)")
	findIdentityCmd.Flags().StringP("string", "s", "", "Specify optional policy-specific string")
	findIdentityCmd.Flags().BoolP("valid-only", "v", false, "Show valid identities only (default: show all)")
	rootCmd.AddCommand(findIdentityCmd)
	carapace.Gen(findIdentityCmd).FlagCompletion(carapace.ActionMap{
		"policy": carapace.ActionValues(
			"basic",
			"ssl-client",
			"ssl-server",
			"smime",
			"eap",
			"ipsec",
			"ichat",
			"codesigning",
			"sys-default",
			"sys-kerberos-kdc",
		),
	})
}
