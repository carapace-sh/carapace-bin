package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_credentials_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete EC2 credentials",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_credentials_deleteCmd).Standalone()

	ec2_credentials_deleteCmd.Flags().String("user", "", "Delete credentials for user (name or ID)")
	ec2_credentials_deleteCmd.Flags().String("user-domain", "", "Domain the user belongs to (name or ID).")
	ec2_credentialsCmd.AddCommand(ec2_credentials_deleteCmd)
}
