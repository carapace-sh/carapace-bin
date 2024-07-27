package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var envEncryptCmd = &cobra.Command{
	Use:   "env:encrypt [--key [KEY]] [--cipher [CIPHER]] [--env [ENV]] [--prune] [--force]",
	Short: "Encrypt an environment file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(envEncryptCmd).Standalone()

	envEncryptCmd.Flags().String("cipher", "", "The encryption cipher")
	envEncryptCmd.Flags().String("env", "", "The environment to be encrypted")
	envEncryptCmd.Flags().Bool("force", false, "Overwrite the existing encrypted environment file")
	envEncryptCmd.Flags().String("key", "", "The encryption key")
	envEncryptCmd.Flags().Bool("prune", false, "Delete the original environment file")
	rootCmd.AddCommand(envEncryptCmd)
}
