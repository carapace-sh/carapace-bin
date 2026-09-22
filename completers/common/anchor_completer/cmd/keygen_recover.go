package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var keygen_recoverCmd = &cobra.Command{
	Use:   "recover",
	Short: "Recover a keypair from a seed phrase",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(keygen_recoverCmd).Standalone()

	keygen_recoverCmd.Flags().BoolP("force", "f", false, "Overwrite the output file if it exists")
	keygen_recoverCmd.Flags().BoolP("help", "h", false, "Print help")
	keygen_recoverCmd.Flags().Bool("no-passphrase", false, "Do not prompt for a passphrase")
	keygen_recoverCmd.Flags().StringP("outfile", "o", "", "Path to recovered keypair file")
	keygen_recoverCmd.Flags().Bool("skip-seed-phrase-validation", false, "Skip seed phrase validation")
	keygenCmd.AddCommand(keygen_recoverCmd)

	carapace.Gen(keygen_recoverCmd).FlagCompletion(carapace.ActionMap{
		"outfile": carapace.ActionFiles(),
	})
}
