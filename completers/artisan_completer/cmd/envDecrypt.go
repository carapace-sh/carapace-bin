package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var envDecryptCmd = &cobra.Command{
	Use:   "env:decrypt [--key [KEY]] [--cipher [CIPHER]] [--env [ENV]] [--force] [--path [PATH]] [--filename [FILENAME]]",
	Short: "Decrypt an environment file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(envDecryptCmd).Standalone()

	envDecryptCmd.Flags().String("cipher", "", "The encryption cipher")
	envDecryptCmd.Flags().String("env", "", "The environment to be decrypted")
	envDecryptCmd.Flags().String("filename", "", "Filename of the decrypted file")
	envDecryptCmd.Flags().Bool("force", false, "Overwrite the existing environment file")
	envDecryptCmd.Flags().String("key", "", "The encryption key")
	envDecryptCmd.Flags().String("path", "", "Path to write the decrypted file")
	rootCmd.AddCommand(envDecryptCmd)
}
