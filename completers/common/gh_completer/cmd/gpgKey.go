package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gpgKeyCmd = &cobra.Command{
	Use:   "gpg-key <command>",
	Short: "Manage GPG keys",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gpgKeyCmd).Standalone()

	rootCmd.AddCommand(gpgKeyCmd)
}
