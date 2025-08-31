package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openssl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var fipsinstallCmd = &cobra.Command{
	Use:     "fipsinstall",
	Short:   "FIPS configuration installation",
	GroupID: "standard",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fipsinstallCmd).Standalone()

	fipsinstallCmd.Flags().StringS("config", "config", "", "The parent config to verify")
	fipsinstallCmd.Flags().StringS("corrupt_desc", "corrupt_desc", "", "Corrupt a self test by description")
	fipsinstallCmd.Flags().StringS("corrupt_type", "corrupt_type", "", "Corrupt a self test by type")
	fipsinstallCmd.Flags().BoolS("dsa_sign_disabled", "dsa_sign_disabled", false, "Disallow DSA signing")
	fipsinstallCmd.Flags().BoolS("ecdh_cofactor_check", "ecdh_cofactor_check", false, "Enable Cofactor check for ECDH")
	fipsinstallCmd.Flags().BoolS("ems_check", "ems_check", false, "Enable the run-time FIPS check for EMS during TLS1_PRF")
	fipsinstallCmd.Flags().BoolS("hkdf_digest_check", "hkdf_digest_check", false, "Enable digest check for HKDF")
	fipsinstallCmd.Flags().BoolS("hkdf_key_check", "hkdf_key_check", false, "Enable key check for HKDF")
	fipsinstallCmd.Flags().BoolS("hmac_key_check", "hmac_key_check", false, "Enable key check for HMAC")
	fipsinstallCmd.Flags().StringS("in", "in", "", "Input config file, used when verifying")
	fipsinstallCmd.Flags().BoolS("kbkdf_key_check", "kbkdf_key_check", false, "Enable key check for KBKDF")
	fipsinstallCmd.Flags().BoolS("kmac_key_check", "kmac_key_check", false, "Enable key check for KMAC")
	fipsinstallCmd.Flags().StringS("mac_name", "mac_name", "", "MAC name")
	fipsinstallCmd.Flags().StringSliceS("macopt", "macopt", nil, "MAC algorithm parameters in n:v form. See 'PARAMETER NAMES' in the EVP_MAC_ docs")
	fipsinstallCmd.Flags().StringS("module", "module", "", "File name of the provider module")
	fipsinstallCmd.Flags().BoolS("no_conditional_errors", "no_conditional_errors", false, "Disable the ability of the fips module to enter an error state if any conditional self tests fail")
	fipsinstallCmd.Flags().BoolS("no_drbg_truncated_digests", "no_drbg_truncated_digests", false, "Disallow truncated digests with Hash and HMAC DRBGs")
	fipsinstallCmd.Flags().BoolS("no_pbkdf2_lower_bound_check", "no_pbkdf2_lower_bound_check", false, "Disable lower bound check for PBKDF2")
	fipsinstallCmd.Flags().BoolS("no_security_checks", "no_security_checks", false, "Disable the run-time FIPS security checks in the module")
	fipsinstallCmd.Flags().BoolS("no_short_mac", "no_short_mac", false, "Disallow short MAC output")
	fipsinstallCmd.Flags().BoolS("noout", "noout", false, "Disable logging of self test events")
	fipsinstallCmd.Flags().StringS("out", "out", "", "Output config file, used when generating")
	fipsinstallCmd.Flags().BoolS("pedantic", "pedantic", false, "Set options for strict FIPS compliance")
	fipsinstallCmd.Flags().StringS("provider_name", "provider_name", "", "FIPS provider name")
	fipsinstallCmd.Flags().BoolS("quiet", "quiet", false, "No messages, just exit status")
	fipsinstallCmd.Flags().BoolS("rsa_pkcs15_padding_disabled", "rsa_pkcs15_padding_disabled", false, "Disallow PKCS#1 version 1.5 padding for RSA encryption")
	fipsinstallCmd.Flags().BoolS("rsa_pss_saltlen_check", "rsa_pss_saltlen_check", false, "Enable salt length check for RSA-PSS signature operations")
	fipsinstallCmd.Flags().BoolS("rsa_sign_x931_disabled", "rsa_sign_x931_disabled", false, "Disallow X931 Padding for RSA signing")
	fipsinstallCmd.Flags().StringS("section_name", "section_name", "", "FIPS Provider config section name (optional)")
	fipsinstallCmd.Flags().BoolS("self_test_oninstall", "self_test_oninstall", false, "Forces self tests to run once on module installation")
	fipsinstallCmd.Flags().BoolS("self_test_onload", "self_test_onload", false, "Forces self tests to always run on module load")
	fipsinstallCmd.Flags().BoolS("signature_digest_check", "signature_digest_check", false, "Enable checking for approved digests for signatures")
	fipsinstallCmd.Flags().BoolS("sshkdf_digest_check", "sshkdf_digest_check", false, "Enable digest check for SSHKDF")
	fipsinstallCmd.Flags().BoolS("sshkdf_key_check", "sshkdf_key_check", false, "Enable key check for SSHKDF")
	fipsinstallCmd.Flags().BoolS("sskdf_digest_check", "sskdf_digest_check", false, "Enable digest check for SSKDF")
	fipsinstallCmd.Flags().BoolS("sskdf_key_check", "sskdf_key_check", false, "Enable key check for SSKDF")
	fipsinstallCmd.Flags().BoolS("tdes_encrypt_disabled", "tdes_encrypt_disabled", false, "Disallow Triple-DES encryption")
	fipsinstallCmd.Flags().BoolS("tls13_kdf_digest_check", "tls13_kdf_digest_check", false, "Enable digest check for TLS13-KDF")
	fipsinstallCmd.Flags().BoolS("tls13_kdf_key_check", "tls13_kdf_key_check", false, "Enable key check for TLS13-KDF")
	fipsinstallCmd.Flags().BoolS("tls1_prf_digest_check", "tls1_prf_digest_check", false, "Enable digest check for TLS1-PRF")
	fipsinstallCmd.Flags().BoolS("tls1_prf_key_check", "tls1_prf_key_check", false, "Enable key check for TLS1-PRF")
	fipsinstallCmd.Flags().BoolS("verify", "verify", false, "Verify a config file instead of generating one")
	fipsinstallCmd.Flags().BoolS("x942kdf_key_check", "x942kdf_key_check", false, "Enable key check for X942KDF")
	fipsinstallCmd.Flags().BoolS("x963kdf_digest_check", "x963kdf_digest_check", false, "Enable digest check for X963KDF")
	fipsinstallCmd.Flags().BoolS("x963kdf_key_check", "x963kdf_key_check", false, "Enable key check for X963KDF")
	rootCmd.AddCommand(fipsinstallCmd)

	carapace.Gen(fipsinstallCmd).FlagCompletion(carapace.ActionMap{
		"config": carapace.ActionFiles(),
		"in":     carapace.ActionFiles(),
		"macopt": carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return carapace.ActionValues(
					"key",
					"hexkey",
					"digest",
				).Suffix(":")
			case 1:
				switch c.Parts[0] {
				case "digest":
					return action.ActionDigestAlgorithms(fipsinstallCmd)
				default:
					return carapace.ActionValues()
				}
			default:
				return carapace.ActionValues()
			}
		}),
		"module": carapace.ActionFiles(),
		"out":    carapace.ActionFiles(),
	})
}
