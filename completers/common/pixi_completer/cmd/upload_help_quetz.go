package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var upload_help_quetzCmd = &cobra.Command{
	Use:   "quetz",
	Short: "Upload to a Quetz server. Authentication is used from the keychain / auth-file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(upload_help_quetzCmd).Standalone()

	upload_helpCmd.AddCommand(upload_help_quetzCmd)
}
