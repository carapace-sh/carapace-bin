package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sshKey_addCmd = &cobra.Command{
	Use:   "add [key-file]",
	Short: "Add an SSH key to your GitLab account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sshKey_addCmd).Standalone()

	sshKey_addCmd.Flags().StringP("expires-at", "e", "", "The expiration date of the SSH key. Uses ISO 8601 format: YYYY-MM-DDTHH:MM:SSZ.")
	sshKey_addCmd.Flags().StringP("title", "t", "", "New SSH key's title.")
	sshKey_addCmd.Flags().StringP("usage-type", "u", "", "Usage scope for the key. Possible values: 'auth', 'signing' or 'auth_and_signing'. Default value: 'auth_and_signing'.")
	sshKey_addCmd.MarkFlagRequired("title")
	sshKeyCmd.AddCommand(sshKey_addCmd)

	carapace.Gen(sshKey_addCmd).FlagCompletion(carapace.ActionMap{
		"usage-type": carapace.ActionValues("auth", "signing", "auth_and_signing"),
	})

	carapace.Gen(sshKey_addCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
