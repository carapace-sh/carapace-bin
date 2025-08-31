package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "openssl",
	Short: "OpenSSL command line program",
	Long:  "https://www.openssl.org/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.AddGroup(
		&cobra.Group{ID: "digest", Title: "Digest Commands"},
		&cobra.Group{ID: "cipher", Title: "Cipher Commands"},
		&cobra.Group{ID: "standard", Title: "Standard Commands"},
	)

	rootCmd.PersistentFlags().BoolS("help", "help", false, "Display this summary")
}
