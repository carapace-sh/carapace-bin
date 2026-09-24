package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var credential_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set credential properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(credential_setCmd).Standalone()

	credential_setCmd.Flags().String("data", "", "New credential data")
	credential_setCmd.Flags().String("project", "", "Project which limits the scope of the credential (name or ID)")
	credential_setCmd.Flags().String("type", "", "New credential type: cert, ec2, totp and so on")
	credential_setCmd.Flags().String("user", "", "User that owns the credential (name or ID)")
	credential_setCmd.MarkFlagRequired("data")
	credential_setCmd.MarkFlagRequired("type")
	credential_setCmd.MarkFlagRequired("user")
	credentialCmd.AddCommand(credential_setCmd)
}
