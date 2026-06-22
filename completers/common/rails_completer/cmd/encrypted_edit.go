package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var encryptedEditCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit an encrypted file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(encryptedEditCmd).Standalone()
	encryptedEditCmd.Flags().BoolP("help", "h", false, "Show help")
	encryptedEditCmd.Flags().StringP("key", "k", "config/master.key", "Path to encryption key")

	carapace.Gen(encryptedEditCmd).PositionalCompletion(
		carapace.ActionFiles("*.yml.enc", "*.yaml.enc", "*.json.enc"),
	)
	carapace.Gen(encryptedEditCmd).FlagCompletion(carapace.ActionMap{
		"key": carapace.ActionFiles(),
	})
}
