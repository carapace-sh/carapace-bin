package cmd

import (
	"github.com/spf13/cobra"
)

var gpgKeyCmd = &cobra.Command{
	Use:   "gpg-key",
	Short: "Manage GPG keys",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	rootCmd.AddCommand(gpgKeyCmd)
}
