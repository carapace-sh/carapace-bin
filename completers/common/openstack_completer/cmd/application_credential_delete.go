package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var application_credential_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete application credentials(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(application_credential_deleteCmd).Standalone()

	application_credentialCmd.AddCommand(application_credential_deleteCmd)
}
