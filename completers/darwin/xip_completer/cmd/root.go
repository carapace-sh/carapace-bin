package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "xip",
	Short: "create or expand a secure archive",
	Long:  "https://keith.github.io/xcode-manpages/xip.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().String("expand", "", "Expand the archive into the current working directory")
	rootCmd.Flags().String("keychain", "", "Specify a specific keychain to search for the signing identity")
	rootCmd.Flags().String("sign", "", "The name of the identity to use for signing")
	rootCmd.Flags().Bool("timestamp", false, "Include a trusted timestamp with the signature")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"expand":   carapace.ActionFiles(),
		"keychain": carapace.ActionFiles(),
	})
}
