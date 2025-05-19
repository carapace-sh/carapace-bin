package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sshKey_addCmd = &cobra.Command{
	Use:   "add [<key-file>]",
	Short: "Add an SSH key to your GitHub account",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sshKey_addCmd).Standalone()

	sshKey_addCmd.Flags().StringP("title", "t", "", "Title for the new key")
	sshKey_addCmd.Flags().String("type", "", "Type of the ssh key: {authentication|signing}")
	sshKeyCmd.AddCommand(sshKey_addCmd)

	carapace.Gen(sshKey_addCmd).FlagCompletion(carapace.ActionMap{
		"type": carapace.ActionValues("authentication", "signing"),
	})

	carapace.Gen(sshKey_addCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
