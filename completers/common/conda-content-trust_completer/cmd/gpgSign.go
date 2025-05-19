package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gpgSignCmd = &cobra.Command{
	Use:   "gpg-sign",
	Short: "Sign a given piece of metadata usign GPG",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gpgSignCmd).Standalone()

	gpgSignCmd.Flags().BoolP("help", "h", false, "show this help message and exit")
	rootCmd.AddCommand(gpgSignCmd)

	carapace.Gen(gpgSignCmd).PositionalCompletion(
		carapace.ActionValues(),
		carapace.ActionFiles(),
	)
}
