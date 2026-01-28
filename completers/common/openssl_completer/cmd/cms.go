package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openssl_completer/cmd/common"
	"github.com/spf13/cobra"
)

var cmsCmd = &cobra.Command{
	Use:     "cms",
	Short:   "CMS (Cryptographic Message Syntax) command",
	GroupID: "standard",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cmsCmd).Standalone()

	cmsCmd.Flags().StringS("CAfile", "CAfile", "", "Trusted certificates file")
	cmsCmd.Flags().StringS("CApath", "CApath", "", "Trusted certificates directory")
	cmsCmd.Flags().StringS("CAstore", "CAstore", "", "Trusted certificates store URI")
	cmsCmd.Flags().BoolS("EncryptedData_decrypt", "EncryptedData_decrypt", false, "Decrypt CMS \"EncryptedData\" object using symmetric key")
	cmsCmd.Flags().BoolS("EncryptedData_encrypt", "EncryptedData_encrypt", false, "Create CMS \"EncryptedData\" object using symmetric key")
	cmsCmd.Flags().BoolS("aes128-wrap", "aes128-wrap", false, "Use AES128 to wrap key")
	cmsCmd.Flags().BoolS("aes192-wrap", "aes192-wrap", false, "Use AES192 to wrap key")
	cmsCmd.Flags().BoolS("aes256-wrap", "aes256-wrap", false, "Use AES256 to wrap key")
	cmsCmd.Flags().BoolS("asciicrlf", "asciicrlf", false, "Perform CRLF canonicalisation when signing")
	cmsCmd.Flags().BoolS("binary", "binary", false, "Treat input as binary: do not translate to canonical form")
	cmsCmd.Flags().BoolS("cades", "cades", false, "Check/Include signingCertificate attribute (CAdES-BES)")
	cmsCmd.Flags().StringSliceS("cert", "cert", nil, "Recipient certs (optional; used only when encrypting)")
	cmsCmd.Flags().StringS("certfile", "certfile", "", "Extra signer and intermediate CA certificates to include when signing or to use as preferred signer certs and for chain building when verifying")
	cmsCmd.Flags().StringS("certsout", "certsout", "", "Certificate output file")
	cmsCmd.Flags().BoolS("cmsout", "cmsout", false, "Output CMS structure")
	cmsCmd.Flags().BoolS("compress", "compress", false, "Create a CMS \"CompressedData\" object")
	cmsCmd.Flags().StringS("config", "config", "", "Load a configuration file (this may load modules)")
	cmsCmd.Flags().StringS("content", "content", "", "Supply or override content for detached signature")
	cmsCmd.Flags().BoolS("crlfeol", "crlfeol", false, "Use CRLF as EOL termination instead of LF only")
	cmsCmd.Flags().BoolS("data_create", "data_create", false, "Create a CMS \"Data\" object")
	cmsCmd.Flags().BoolS("data_out", "data_out", false, "Copy CMS \"Data\" object to output")
	cmsCmd.Flags().BoolS("debug_decrypt", "debug_decrypt", false, "Disable MMA protection, return error if no recipient found (see doc)")
	cmsCmd.Flags().BoolS("decrypt", "decrypt", false, "Decrypt encrypted message")
	cmsCmd.Flags().BoolS("des3-wrap", "des3-wrap", false, "Use 3DES-EDE to wrap key")
	cmsCmd.Flags().StringS("digest", "digest", "", "Sign a pre-computed digest in hex notation")
	cmsCmd.Flags().BoolS("digest_create", "digest_create", false, "Create a CMS \"DigestedData\" object")
	cmsCmd.Flags().BoolS("digest_verify", "digest_verify", false, "Verify a CMS \"DigestedData\" object and output it")
	cmsCmd.Flags().StringS("econtent_type", "econtent_type", "", "OID for external content")
	cmsCmd.Flags().BoolS("encrypt", "encrypt", false, "Encrypt message")
	cmsCmd.Flags().StringS("from", "from", "", "From address")
	cmsCmd.Flags().StringS("in", "in", "", "Input file")
	cmsCmd.Flags().BoolS("indef", "indef", false, "Same as -stream")
	cmsCmd.Flags().StringS("inform", "inform", "", "Input format SMIME (default), PEM or DER")
	cmsCmd.Flags().StringS("inkey", "inkey", "", "Input private key (if not signer or recipient)")
	cmsCmd.Flags().StringS("kekcipher", "kekcipher", "", "The key encryption algorithm to use")
	cmsCmd.Flags().StringS("keyform", "keyform", "", "Input private key format (DER/PEM)")
	cmsCmd.Flags().BoolS("keyid", "keyid", false, "Use subject key identifier")
	cmsCmd.Flags().StringSliceS("keyopt", "keyopt", nil, "Set public key parameters as n:v pairs")
	cmsCmd.Flags().StringS("md", "md", "", "Digest algorithm to use")
	cmsCmd.Flags().StringSliceS("nameopt", "nameopt", nil, "For the -print option specifies various strings printing options")
	cmsCmd.Flags().BoolS("no-CAfile", "no-CAfile", false, "Do not load the default certificates file")
	cmsCmd.Flags().BoolS("no-CApath", "no-CApath", false, "Do not load certificates from the default certificates directory")
	cmsCmd.Flags().BoolS("no-CAstore", "no-CAstore", false, "Do not load certificates from the default certificates store")
	cmsCmd.Flags().BoolS("no_attr_verify", "no_attr_verify", false, "Do not verify signed attribute signatures")
	cmsCmd.Flags().BoolS("no_content_verify", "no_content_verify", false, "Do not verify signed content signatures")
	cmsCmd.Flags().BoolS("no_signing_time", "no_signing_time", false, "Omit the signing time attribute")
	cmsCmd.Flags().BoolS("noattr", "noattr", false, "Don't include any signed attributes")
	cmsCmd.Flags().BoolS("nocerts", "nocerts", false, "Don't include signer's certificate when signing")
	cmsCmd.Flags().BoolS("nodetach", "nodetach", false, "Use opaque signing")
	cmsCmd.Flags().BoolS("noindef", "noindef", false, "Disable CMS streaming")
	cmsCmd.Flags().BoolS("nointern", "nointern", false, "Don't search certificates in message for signer")
	cmsCmd.Flags().BoolS("noout", "noout", false, "For the -cmsout operation do not output the parsed CMS structure")
	cmsCmd.Flags().BoolS("nosigs", "nosigs", false, "Don't verify message signature")
	cmsCmd.Flags().BoolS("nosmimecap", "nosmimecap", false, "Omit the SMIMECapabilities attribute")
	cmsCmd.Flags().BoolS("noverify", "noverify", false, "Don't verify signers certificate")
	cmsCmd.Flags().StringS("originator", "originator", "", "Originator certificate file")
	cmsCmd.Flags().StringS("out", "out", "", "Output file")
	cmsCmd.Flags().StringS("outform", "outform", "", "Output format SMIME (default), PEM or DER")
	cmsCmd.Flags().StringS("passin", "passin", "", "Input file pass phrase source")
	cmsCmd.Flags().BoolS("print", "print", false, "For the -cmsout operation print out all fields of the CMS structure")
	cmsCmd.Flags().StringS("pwri_password", "pwri_password", "", "Specific password for recipient")
	cmsCmd.Flags().StringS("rctform", "rctform", "", "Receipt file format")
	cmsCmd.Flags().BoolS("receipt_request_all", "receipt_request_all", false, "When signing, create a receipt request for all recipients")
	cmsCmd.Flags().BoolS("receipt_request_first", "receipt_request_first", false, "When signing, create a receipt request for first recipient")
	cmsCmd.Flags().StringS("receipt_request_from", "receipt_request_from", "", "Create signed receipt request with specified email address")
	cmsCmd.Flags().BoolS("receipt_request_print", "receipt_request_print", false, "Print CMS Receipt Request")
	cmsCmd.Flags().StringS("receipt_request_to", "receipt_request_to", "", "Create signed receipt targeted to specified address")
	cmsCmd.Flags().StringS("recip", "recip", "", "Recipient cert file")
	cmsCmd.Flags().StringS("recip_kdf", "recip_kdf", "", "Set KEMRecipientInfo KDF for current recipient")
	cmsCmd.Flags().StringS("recip_ukm", "recip_ukm", "", "KEMRecipientInfo user keying material for current recipient, in hex notation")
	cmsCmd.Flags().BoolS("resign", "resign", false, "Resign a signed message")
	cmsCmd.Flags().StringS("secretkey", "secretkey", "", "Use specified hex-encoded key to decrypt/encrypt recipients or content")
	cmsCmd.Flags().StringS("secretkeyid", "secretkeyid", "", "Identity of the -secretkey for CMS \"KEKRecipientInfo\" object")
	cmsCmd.Flags().BoolS("sign", "sign", false, "Sign message")
	cmsCmd.Flags().BoolS("sign_receipt", "sign_receipt", false, "Generate a signed receipt for a message")
	cmsCmd.Flags().StringS("signer", "signer", "", "Signer certificate(s) input/output file")
	cmsCmd.Flags().BoolS("stream", "stream", false, "Enable CMS streaming")
	cmsCmd.Flags().StringS("subject", "subject", "", "Subject")
	cmsCmd.Flags().BoolS("text", "text", false, "Include or delete text MIME headers")
	cmsCmd.Flags().StringS("to", "to", "", "To address")
	cmsCmd.Flags().BoolS("uncompress", "uncompress", false, "Uncompress a CMS \"CompressedData\" object")
	cmsCmd.Flags().BoolS("verify", "verify", false, "Verify signed message")
	cmsCmd.Flags().StringS("verify_receipt", "verify_receipt", "", "Verify receipts; exit if receipt signatures do not verify")
	cmsCmd.Flags().BoolS("verify_retcode", "verify_retcode", false, "Exit non-zero on verification failure")
	cmsCmd.Flags().StringS("wrap", "wrap", "", "Key wrap algorithm to use when encrypting with key agreement or key encapsulation")
	common.AddProviderFlags(cmsCmd)
	common.AddRandomStateFlags(cmsCmd)
	common.AddValidationFlags(cmsCmd)
	rootCmd.AddCommand(cmsCmd)

	carapace.Gen(cmsCmd).FlagCompletion(carapace.ActionMap{
		"CAfile":         carapace.ActionFiles(),
		"CApath":         carapace.ActionDirectories(),
		"cert":           carapace.ActionFiles(),
		"certfile":       carapace.ActionFiles(),
		"certsout":       carapace.ActionFiles(),
		"config":         carapace.ActionFiles(),
		"content":        carapace.ActionFiles(),
		"in":             carapace.ActionFiles(),
		"inform":         carapace.ActionValues("DER", "PEM", "SMIME"),
		"inkey":          carapace.ActionFiles(),
		"keyform":        carapace.ActionValues("DER", "PEM", "P12"),
		"originator":     carapace.ActionFiles(),
		"out":            carapace.ActionFiles(),
		"outform":        carapace.ActionValues("DER", "PEM", "SMIME"),
		"rctform":        carapace.ActionValues("DER", "PEM", "SMIME"),
		"recip":          carapace.ActionFiles(),
		"signer":         carapace.ActionFiles(),
		"verify_receipt": carapace.ActionFiles(),
	})
}
