package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openssl_completer/cmd/common"
	"github.com/spf13/cobra"
)

var speedCmd = &cobra.Command{
	Use:     "speed",
	Short:   "Algorithm Speed Measurement",
	GroupID: "standard",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(speedCmd).Standalone()

	speedCmd.Flags().BoolS("aead", "aead", false, "Benchmark EVP-named AEAD cipher in TLS-like sequence")
	speedCmd.Flags().StringS("async_jobs", "async_jobs", "", "Enable async mode and start specified number of jobs")
	speedCmd.Flags().StringS("bytes", "bytes", "", "Run [non-PKI] benchmarks on custom-sized buffer")
	speedCmd.Flags().StringS("cmac", "cmac", "", "CMAC using EVP-named cipher")
	speedCmd.Flags().StringS("config", "config", "", "Load a configuration file (this may load modules)")
	speedCmd.Flags().BoolS("decrypt", "decrypt", false, "Time decryption instead of encryption (only EVP)")
	speedCmd.Flags().BoolS("elapsed", "elapsed", false, "Use wall-clock time instead of CPU user time as divisor")
	speedCmd.Flags().StringS("evp", "evp", "", "Use EVP-named cipher or digest")
	speedCmd.Flags().StringS("hmac", "hmac", "", "HMAC using EVP-named digest")
	speedCmd.Flags().BoolS("kem-algorithms", "kem-algorithms", false, "Benchmark KEM algorithms")
	speedCmd.Flags().BoolS("mb", "mb", false, "Enable (tls1>=1) multi-block mode on EVP-named cipher")
	speedCmd.Flags().StringS("misalign", "misalign", "", "Use specified offset to mis-align buffers")
	speedCmd.Flags().BoolS("mlock", "mlock", false, "Lock memory for better result determinism")
	speedCmd.Flags().BoolS("mr", "mr", false, "Produce machine readable output")
	speedCmd.Flags().StringS("multi", "multi", "", "Run benchmarks in parallel")
	speedCmd.Flags().StringS("primes", "primes", "", "Specify number of primes (for RSA only)")
	speedCmd.Flags().StringS("seconds", "seconds", "", "Run benchmarks for specified amount of seconds")
	speedCmd.Flags().BoolS("signature-algorithms", "signature-algorithms", false, "Benchmark signature algorithms")
	speedCmd.Flags().BoolS("testmode", "testmode", false, "Run the speed command in test mode")
	common.AddProviderFlags(speedCmd)
	common.AddRandomStateFlags(speedCmd)
	rootCmd.AddCommand(speedCmd)

	carapace.Gen(speedCmd).FlagCompletion(carapace.ActionMap{
		"config": carapace.ActionFiles(),
	})
}
