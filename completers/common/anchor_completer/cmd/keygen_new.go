package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var keygen_newCmd = &cobra.Command{
	Use:   "new",
	Short: "Generate a new keypair",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(keygen_newCmd).Standalone()

	keygen_newCmd.Flags().BoolP("force", "f", false, "Overwrite the output file if it exists")
	keygen_newCmd.Flags().BoolP("help", "h", false, "Print help")
	keygen_newCmd.Flags().Bool("no-passphrase", false, "Do not prompt for a passphrase")
	keygen_newCmd.Flags().StringP("outfile", "o", "", "Path to generated keypair file")
	keygen_newCmd.Flags().Bool("silent", false, "Do not display the generated pubkey")
	keygen_newCmd.Flags().StringP("word-count", "w", "12", "Number of words in the mnemonic phrase [possible values: 12, 15, 18, 21, 24]")
	keygenCmd.AddCommand(keygen_newCmd)
}
