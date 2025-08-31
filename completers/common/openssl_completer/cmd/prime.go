package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openssl_completer/cmd/common"
	"github.com/spf13/cobra"
)

var primeCmd = &cobra.Command{
	Use:     "prime",
	Short:   "Compute prime numbers",
	GroupID: "standard",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(primeCmd).Standalone()

	primeCmd.Flags().StringS("bits", "bits", "", "Size of number in bits")
	primeCmd.Flags().StringS("checks", "checks", "", "Number of checks")
	primeCmd.Flags().BoolS("generate", "generate", false, "Generate a prime")
	primeCmd.Flags().BoolS("hex", "hex", false, "Enables hex format for output from prime generation or input to primality checking")
	primeCmd.Flags().BoolS("in", "in", false, "Provide file names containing numbers for primality checking")
	primeCmd.Flags().BoolS("safe", "safe", false, "When used with -generate, generate a safe prime")
	common.AddProviderFlags(primeCmd)
	rootCmd.AddCommand(primeCmd)
}
