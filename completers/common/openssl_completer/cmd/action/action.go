package action

import (
	"regexp"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

// ActionDigestAlgorithms completes message digest algorithms
//
//	BLAKE2b512
//	BLAKE2B-512
func ActionDigestAlgorithms(cmd *cobra.Command) carapace.Action {
	return actionOpenSSLAlgorithms(cmd, "-digest-algorithms")
}

// ActionCipherAlgorithms completes symmetric cipher algorithms
//
//	AES128
//	AES128-WRAP
func ActionCipherAlgorithms(cmd *cobra.Command) carapace.Action {
	return actionOpenSSLAlgorithms(cmd, "-cipher-algorithms")
}

// ActionKDFAlgorithms completes key derivation and pseudo random function algorithms
//
//	ARGON2D
//	ARGON2I
func ActionKDFAlgorithms(cmd *cobra.Command) carapace.Action {
	return actionOpenSSLAlgorithms(cmd, "-kdf-algorithms")
}

// ActionMACAlgorithms completes message authentication code algorithms
//
//	BLAKE2BMAC
//	BLAKE2SMAC
func ActionMACAlgorithms(cmd *cobra.Command) carapace.Action {
	return actionOpenSSLAlgorithms(cmd, "-mac-algorithms")
}

// ActionKEMAlgorithms completes key encapsulation mechanism algorithms
//
//	EC
//	id-alg-ml-kem-512
func ActionKEMAlgorithms(cmd *cobra.Command) carapace.Action {
	return actionOpenSSLAlgorithms(cmd, "-kem-algorithms")
}

// ActionKeyExchangeAlgorithms completes key exchange algorithms
//
//	DH
//	dhKeyAgreement
func ActionKeyExchangeAlgorithms(cmd *cobra.Command) carapace.Action {
	return actionOpenSSLAlgorithms(cmd, "-key-exchange-algorithms")
}

func actionOpenSSLAlgorithms(cmd *cobra.Command, flag string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		providers, _ := cmd.Flags().GetStringSlice("provider")
		if len(providers) == 0 {
			providers = []string{"default"}
		}

		args := []string{"list", flag}
		for _, provider := range providers {
			args = append(args, "-provider", provider)
		}

		return carapace.ActionExecCommand("openssl", args...)(func(output []byte) carapace.Action {
			re := regexp.MustCompile(`[[:alpha:]][[:alnum:]_/-]*`)
			lines := strings.Split(string(output), "\n")

			var values []string
			for _, line := range lines {
				if idx := strings.Index(line, "@"); idx != -1 {
					line = strings.TrimSpace(line[:idx])
					matches := re.FindAllString(line, -1)
					values = append(values, matches...)
				}
			}
			return carapace.ActionValues(values...)
		})
	})
}
