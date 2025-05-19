package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var auth_gitCredentialCmd = &cobra.Command{
	Use:    "git-credential",
	Short:  "Implements git credential helper protocol",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auth_gitCredentialCmd).Standalone()

	authCmd.AddCommand(auth_gitCredentialCmd)
}
