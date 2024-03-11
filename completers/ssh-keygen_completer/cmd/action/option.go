package action

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
)

func ActionOptions() carapace.Action {
	return carapace.ActionMultiParts("=", func(c carapace.Context) carapace.Action {
		switch len(c.Parts) {
		case 0:
			return carapace.ActionValuesDescribed(
				"application", "Override the default FIDO application/origin string of `“ssh:`”",
				"challenge", "Specifies a path to a challenge string that will be passed to the FIDO token during key generation",
				"device", "Explicitly specify a fido device to use",
				"no-touch-required", "Indicate that the generated private key should not require touch events when making signatures",
				"resident", "Indicate that the key should be stored on the FIDO authenticator itself",
				"user", "A username to be associated with a resident key",
				"verify-required", "Indicate that this private key should require user verification for each signature",
				"write-attestation", "Record the attestation data returned from FIDO tokens during key generation",
				"hashalg", "Selects the hash algorithm to use for hashing the message to be signed",
				"print-pubkey", "Print the full public key to standard output after signature verification",
				"verify-time", "Specifies a time to use when validating signatures instead of the current time",
			).Invoke(c).Suffix("=").ToA()
		case 1:
			return map[string]carapace.Action{
				"application":       carapace.ActionValues(),
				"challenge":         carapace.ActionValues(),
				"device":            carapace.ActionValues(),
				"no-touch-required": carapace.ActionValues(),
				"resident":          carapace.ActionValues(),
				"user":              os.ActionUsers(),
				"verify-required":   carapace.ActionValues(),
				"write-attestation": carapace.ActionFiles(),
				"hashalg":           carapace.ActionValues("sha256", "sha512"),
				"print-pubkey":      carapace.ActionValues(),
				"verify-time":       carapace.ActionValues(),
			}[c.Parts[0]]
		default:
			return carapace.ActionValues()
		}
	})
}
