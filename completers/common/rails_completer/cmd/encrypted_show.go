package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var encryptedShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show an encrypted file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(encryptedShowCmd).Standalone()
	encryptedShowCmd.Flags().BoolP("help", "h", false, "Show help")
	encryptedShowCmd.Flags().StringP("key", "k", "config/master.key", "Path to encryption key")

	carapace.Gen(encryptedShowCmd).PositionalCompletion(
		carapace.ActionFiles("*.yml.enc", "*.yaml.enc", "*.json.enc"),
	)
	carapace.Gen(encryptedShowCmd).FlagCompletion(carapace.ActionMap{
		"key": carapace.ActionFiles(),
	})
}
