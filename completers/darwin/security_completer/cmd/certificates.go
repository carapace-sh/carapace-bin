package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/security"
	"github.com/spf13/cobra"
)

var authorizeCmd = &cobra.Command{
	Use:   "authorize",
	Short: "Perform authorization operations",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var authorizationdbCmd = &cobra.Command{
	Use:   "authorizationdb",
	Short: "Make changes to the authorization policy database",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var cmsCmd = &cobra.Command{
	Use:   "cms",
	Short: "Create or decode CMS messages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var createDbCmd = &cobra.Command{
	Use:   "create-db",
	Short: "Create a db using the DL",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var createKeypairCmd = &cobra.Command{
	Use:   "create-keypair",
	Short: "Create an asymmetric key pair",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var dumpTrustSettingsCmd = &cobra.Command{
	Use:   "dump-trust-settings",
	Short: "Display contents of trust settings",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var errorCmd = &cobra.Command{
	Use:   "error",
	Short: "Display a descriptive message for error codes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export items from a keychain",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var importCmd = &cobra.Command{
	Use:   "import",
	Short: "Import items into a keychain",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var addTrustedCertCmd = &cobra.Command{
	Use:   "add-trusted-cert",
	Short: "Add trusted certificate(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var removeTrustedCertCmd = &cobra.Command{
	Use:   "remove-trusted-cert",
	Short: "Remove trusted certificate(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var trustSettingsExportCmd = &cobra.Command{
	Use:   "trust-settings-export",
	Short: "Export trust settings",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var trustSettingsImportCmd = &cobra.Command{
	Use:   "trust-settings-import",
	Short: "Import trust settings",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var userTrustSettingsEnableCmd = &cobra.Command{
	Use:   "user-trust-settings-enable",
	Short: "Display or manipulate user-level trust settings",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var verifyCertCmd = &cobra.Command{
	Use:   "verify-cert",
	Short: "Verify certificate(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(authorizeCmd).Standalone()
	authorizeCmd.Flags().BoolP("destroy", "d", false, "Destroy acquired rights")
	authorizeCmd.Flags().BoolP("externalize", "e", false, "Externalize authref to stdout")
	authorizeCmd.Flags().BoolP("internalize", "i", false, "Internalize authref passed on stdin")
	authorizeCmd.Flags().BoolP("least-privileged", "l", false, "Operate in least privileged mode")
	authorizeCmd.Flags().BoolP("partial-rights", "p", false, "Allow returning partial rights")
	authorizeCmd.Flags().BoolP("pre-authorize", "P", false, "Pre-authorize rights only")
	authorizeCmd.Flags().BoolP("user-interaction", "u", false, "Allow user interaction")
	authorizeCmd.Flags().BoolP("wait", "w", false, "Wait while holding AuthorizationRef until stdout is closed")
	rootCmd.AddCommand(authorizeCmd)

	carapace.Gen(authorizationdbCmd).Standalone()
	rootCmd.AddCommand(authorizationdbCmd)
	carapace.Gen(authorizationdbCmd).PositionalCompletion(
		carapace.ActionValues("read", "write", "remove"),
	)

	carapace.Gen(cmsCmd).Standalone()
	cmsCmd.Flags().StringP("cert-usage", "u", "", "Set type of certificate usage (0-11)")
	cmsCmd.Flags().StringP("content", "c", "", "Use this detached content file (decode)")
	cmsCmd.Flags().BoolP("create-encrypted", "C", false, "Create a CMS encrypted message")
	cmsCmd.Flags().BoolP("create-enveloped", "E", false, "Create a CMS enveloped message")
	cmsCmd.Flags().BoolP("create-signed", "S", false, "Create a CMS signed message")
	cmsCmd.Flags().BoolP("decode", "D", false, "Decode a CMS message")
	cmsCmd.Flags().StringP("enc-key-pref", "Y", "", "Include an EncryptionKeyPreference attribute (encode)")
	cmsCmd.Flags().StringP("envelope", "e", "", "Specify envelope file")
	cmsCmd.Flags().StringP("hash-algorithm", "H", "", "Hash algorithm (encode, default: SHA1)")
	cmsCmd.Flags().StringP("headers", "h", "", "Generate email headers with info (decode)")
	cmsCmd.Flags().StringP("input", "i", "", "Use infile as source of data (default: stdin)")
	cmsCmd.Flags().StringP("keychain", "k", "", "Specify keychain to use")
	cmsCmd.Flags().BoolP("no-content", "T", false, "Do not include content in CMS message (encode)")
	cmsCmd.Flags().StringP("output", "o", "", "Use outfile as destination of data (default: stdout)")
	cmsCmd.Flags().StringP("password", "p", "", "Use password as key db password")
	cmsCmd.Flags().StringP("recipients", "r", "", "Create envelope for recipients (encode)")
	cmsCmd.Flags().StringP("signer", "N", "", "Use certificate named for signing (encode)")
	cmsCmd.Flags().BoolP("signing-time", "G", false, "Include a signing time attribute (encode)")
	cmsCmd.Flags().BoolP("single-byte", "s", false, "Pass data a single byte at a time to CMS")
	cmsCmd.Flags().BoolP("smime-capabilities", "P", false, "Include SMIMECapabilities attribute (encode)")
	cmsCmd.Flags().StringP("subject-key-id", "Z", "", "Find a certificate by subject key ID (encode)")
	cmsCmd.Flags().BoolP("suppress-content", "n", false, "Suppress output of content (decode)")
	cmsCmd.Flags().BoolP("verbose", "v", false, "Print debugging information")
	rootCmd.AddCommand(cmsCmd)
	carapace.Gen(cmsCmd).FlagCompletion(carapace.ActionMap{
		"content":        carapace.ActionFiles(),
		"envelope":       carapace.ActionFiles(),
		"hash-algorithm": carapace.ActionValues("MD2", "MD4", "MD5", "SHA1", "SHA256", "SHA384", "SHA512"),
		"input":          carapace.ActionFiles(),
		"keychain":       security.ActionKeychains(),
		"output":         carapace.ActionFiles(),
	})

	carapace.Gen(createDbCmd).Standalone()
	createDbCmd.Flags().BoolP("autocommit-off", "a", false, "Turn off autocommit")
	createDbCmd.Flags().StringP("dl", "g", "", "Use AppleDL (default) or AppleCspDL")
	createDbCmd.Flags().StringP("mode", "m", "", "Set the file permissions to mode")
	createDbCmd.Flags().BoolP("openparams", "o", false, "Force using openparams argument")
	rootCmd.AddCommand(createDbCmd)

	carapace.Gen(createKeypairCmd).Standalone()
	createKeypairCmd.Flags().StringP("algorithm", "a", "", "Algorithm: rsa, dh, dsa, or fee (default: rsa)")
	createKeypairCmd.Flags().BoolP("allow-all", "A", false, "Allow any application to access without warning")
	createKeypairCmd.Flags().StringP("days", "d", "", "Make key valid for specified days from today")
	createKeypairCmd.Flags().StringP("keychain", "k", "", "Use specified keychain instead of default")
	createKeypairCmd.Flags().StringP("size", "s", "", "Keysize in bits (default: 512)")
	createKeypairCmd.Flags().StringP("tool", "T", "", "Specify an application which may access this key")
	createKeypairCmd.Flags().StringP("valid-from", "f", "", "Make key valid from specified date")
	createKeypairCmd.Flags().StringP("valid-to", "t", "", "Make key valid to specified date")
	rootCmd.AddCommand(createKeypairCmd)
	carapace.Gen(createKeypairCmd).FlagCompletion(carapace.ActionMap{
		"algorithm": carapace.ActionValues("rsa", "dh", "dsa", "fee"),
		"keychain":  security.ActionKeychains(),
	})

	carapace.Gen(dumpTrustSettingsCmd).Standalone()
	dumpTrustSettingsCmd.Flags().BoolP("admin", "d", false, "Display trusted admin certs (default: user)")
	dumpTrustSettingsCmd.Flags().BoolP("system", "s", false, "Display trusted system certs (default: user)")
	rootCmd.AddCommand(dumpTrustSettingsCmd)

	carapace.Gen(errorCmd).Standalone()
	rootCmd.AddCommand(errorCmd)

	carapace.Gen(exportCmd).Standalone()
	exportCmd.Flags().StringP("format", "f", "", "Format: openssl, bsafe, pkcs7, pkcs8, pkcs12, x509, etc.")
	exportCmd.Flags().StringP("keychain", "k", "", "Specify keychain from which item(s) will be exported")
	exportCmd.Flags().StringP("output", "o", "", "Write output data to outfile (default: stdout)")
	exportCmd.Flags().StringP("passphrase", "P", "", "Specify wrapping passphrase immediately")
	exportCmd.Flags().BoolP("pem-armour", "p", false, "PEM armour applied to output data")
	exportCmd.Flags().StringP("type", "t", "", "Type: certs, allKeys, pubKeys, privKeys, identities, all")
	exportCmd.Flags().BoolP("wrap", "w", false, "Private keys are to be wrapped on export")
	rootCmd.AddCommand(exportCmd)
	carapace.Gen(exportCmd).FlagCompletion(carapace.ActionMap{
		"format":   carapace.ActionValues("openssl", "bsafe", "pkcs7", "pkcs8", "pkcs12", "x509", "openssh1", "openssh2", "pemseq"),
		"keychain": security.ActionKeychains(),
		"output":   carapace.ActionFiles(),
		"type":     carapace.ActionValues("certs", "allKeys", "pubKeys", "privKeys", "identities", "all"),
	})

	carapace.Gen(importCmd).Standalone()
	importCmd.Flags().BoolP("allow-all", "A", false, "Allow any application to access imported key without warning")
	importCmd.Flags().StringP("attribute", "a", "", "Specify optional extended attribute name and value")
	importCmd.Flags().StringP("format", "f", "", "Format: openssl, bsafe, raw, pkcs7, pkcs8, pkcs12, x509, etc.")
	importCmd.Flags().StringP("keychain", "k", "", "Specify keychain into which item(s) will be imported")
	importCmd.Flags().BoolP("non-extractable", "x", false, "Private keys are non-extractable after being imported")
	importCmd.Flags().StringP("passphrase", "P", "", "Specify unwrapping passphrase immediately")
	importCmd.Flags().StringP("tool", "T", "", "Specify an application which may access imported key")
	importCmd.Flags().StringP("type", "t", "", "Type: cert, pub, priv, session, agg")
	importCmd.Flags().BoolP("wrap", "w", false, "Private keys are wrapped and must be unwrapped on import")
	rootCmd.AddCommand(importCmd)
	carapace.Gen(importCmd).PositionalCompletion(carapace.ActionFiles())
	carapace.Gen(importCmd).FlagCompletion(carapace.ActionMap{
		"format":   carapace.ActionValues("openssl", "bsafe", "raw", "pkcs7", "pkcs8", "pkcs12", "x509", "openssh1", "openssh2", "pemseq"),
		"keychain": security.ActionKeychains(),
		"type":     carapace.ActionValues("cert", "pub", "priv", "session", "agg"),
	})

	carapace.Gen(addTrustedCertCmd).Standalone()
	addTrustedCertCmd.Flags().BoolP("admin", "d", false, "Add to admin cert store (default: user)")
	addTrustedCertCmd.Flags().StringP("allowed-error", "e", "", "Allowed error (integer, certExpired, hostnameMismatch)")
	addTrustedCertCmd.Flags().StringP("app-path", "a", "", "Specify application constraint")
	addTrustedCertCmd.Flags().BoolP("default", "D", false, "Add default setting instead of per-cert setting")
	addTrustedCertCmd.Flags().StringP("key-usage", "u", "", "Specify key usage (integer)")
	addTrustedCertCmd.Flags().StringP("keychain", "k", "", "Specify keychain to which cert is added")
	addTrustedCertCmd.Flags().StringP("policy", "p", "", "Policy constraint: ssl, smime, codeSign, etc.")
	addTrustedCertCmd.Flags().StringP("policy-string", "s", "", "Specify policy-specific string")
	addTrustedCertCmd.Flags().StringP("result-type", "r", "", "trustRoot, trustAsRoot, deny, or unspecified")
	addTrustedCertCmd.Flags().StringP("settings-in", "i", "", "Input trust settings file (default: user domain)")
	addTrustedCertCmd.Flags().StringP("settings-out", "o", "", "Output trust settings file (default: user domain)")
	rootCmd.AddCommand(addTrustedCertCmd)
	carapace.Gen(addTrustedCertCmd).PositionalCompletion(carapace.ActionFiles())
	carapace.Gen(addTrustedCertCmd).FlagCompletion(carapace.ActionMap{
		"keychain":     security.ActionKeychains(),
		"policy":       carapace.ActionValues("ssl", "smime", "codeSign", "IPSec", "iChat", "basic", "swUpdate", "pkgSign", "pkinitClient", "pkinitServer", "eap"),
		"result-type":  carapace.ActionValues("trustRoot", "trustAsRoot", "deny", "unspecified"),
		"settings-in":  carapace.ActionFiles(),
		"settings-out": carapace.ActionFiles(),
	})

	carapace.Gen(removeTrustedCertCmd).Standalone()
	removeTrustedCertCmd.Flags().BoolP("admin", "d", false, "Remove from admin cert store (default: user)")
	removeTrustedCertCmd.Flags().BoolP("default", "D", false, "Remove Default Root Cert setting instead of actual cert setting")
	rootCmd.AddCommand(removeTrustedCertCmd)
	carapace.Gen(removeTrustedCertCmd).PositionalCompletion(carapace.ActionFiles())

	carapace.Gen(trustSettingsExportCmd).Standalone()
	trustSettingsExportCmd.Flags().BoolP("admin", "d", false, "Export admin Trust Settings (default: user)")
	trustSettingsExportCmd.Flags().BoolP("system", "s", false, "Export system Trust Settings (default: user)")
	rootCmd.AddCommand(trustSettingsExportCmd)
	carapace.Gen(trustSettingsExportCmd).PositionalCompletion(carapace.ActionFiles())

	carapace.Gen(trustSettingsImportCmd).Standalone()
	trustSettingsImportCmd.Flags().BoolP("admin", "d", false, "Import admin Trust Settings (default: user)")
	rootCmd.AddCommand(trustSettingsImportCmd)
	carapace.Gen(trustSettingsImportCmd).PositionalCompletion(carapace.ActionFiles())

	carapace.Gen(userTrustSettingsEnableCmd).Standalone()
	userTrustSettingsEnableCmd.Flags().BoolP("disable", "d", false, "Disable user-level Trust Settings")
	userTrustSettingsEnableCmd.Flags().BoolP("enable", "e", false, "Enable user-level Trust Settings")
	rootCmd.AddCommand(userTrustSettingsEnableCmd)

	carapace.Gen(verifyCertCmd).Standalone()
	verifyCertCmd.Flags().StringP("cert", "c", "", "Certificate to verify (DER or PEM)")
	verifyCertCmd.Flags().StringP("email", "e", "", "Specify email address for smime policy")
	verifyCertCmd.Flags().StringP("keychain", "k", "", "Keychain for intermediate certs")
	verifyCertCmd.Flags().BoolP("leaf-is-ca", "l", false, "Leaf certificate is a CA cert")
	verifyCertCmd.Flags().BoolP("local-only", "L", false, "Use local certificates only (no network fetch)")
	verifyCertCmd.Flags().BoolP("no-keychains", "n", false, "Avoid searching any keychains")
	verifyCertCmd.Flags().StringP("policy", "p", "", "Verification policy (default: basic)")
	verifyCertCmd.Flags().BoolP("quiet", "q", false, "Quiet, no stdout or stderr")
	verifyCertCmd.Flags().StringP("root-cert", "r", "", "Root certificate (DER or PEM)")
	verifyCertCmd.Flags().StringP("ssl-host", "s", "", "Specify SSL host name for ssl policy")
	rootCmd.AddCommand(verifyCertCmd)
	carapace.Gen(verifyCertCmd).PositionalCompletion(carapace.ActionFiles())
	carapace.Gen(verifyCertCmd).FlagCompletion(carapace.ActionMap{
		"cert":      carapace.ActionFiles(),
		"keychain":  security.ActionKeychains(),
		"policy":    carapace.ActionValues("ssl", "smime", "codeSign", "IPSec", "iChat", "basic", "swUpdate", "pkgSign", "pkinitClient", "pkinitServer", "eap", "appleID", "macappstore", "timestamping"),
		"root-cert": carapace.ActionFiles(),
	})
}
