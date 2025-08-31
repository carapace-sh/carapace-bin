package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openssl_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/completers/common/openssl_completer/cmd/common"
	"github.com/spf13/cobra"
)

var genrsaCmd = &cobra.Command{
	Use:     "genrsa",
	Short:   "Generation of RSA Private Key",
	GroupID: "standard",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(genrsaCmd).Standalone()

	genrsaCmd.Flags().BoolS("3", "3", false, "(deprecated) Use 3 for the E value")
	genrsaCmd.Flags().BoolS("F4", "F4", false, "Use the Fermat number F4 (0x10001) for the E value")
	genrsaCmd.Flags().StringS("engine", "engine", "", "Use engine, possibly a hardware device")
	genrsaCmd.Flags().BoolS("f4", "f4", false, "Use the Fermat number F4 (0x10001) for the E value")
	genrsaCmd.Flags().StringS("out", "out", "", "Output the key to specified file")
	genrsaCmd.Flags().StringS("passout", "passout", "", "Output file pass phrase source")
	genrsaCmd.Flags().StringS("primes", "primes", "", "Specify number of primes")
	genrsaCmd.Flags().BoolS("quiet", "quiet", false, "Terse output")
	genrsaCmd.Flags().BoolS("traditional", "traditional", false, "Use traditional format for private keys")
	genrsaCmd.Flags().BoolS("verbose", "verbose", false, "Verbose output")
	common.AddProviderFlags(genrsaCmd)
	common.AddRandomStateFlags(genrsaCmd)
	rootCmd.AddCommand(genrsaCmd)

	carapace.Gen(genrsaCmd).FlagCompletion(carapace.ActionMap{
		"engine": action.ActionEngines(),
		"out":    carapace.ActionFiles(),
	})
}
