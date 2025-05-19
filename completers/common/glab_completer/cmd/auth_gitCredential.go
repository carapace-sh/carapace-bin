package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var auth_gitCredentialCmd = &cobra.Command{
	Use:    "git-credential",
	Short:  "Implements Git credential helper manager.",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auth_gitCredentialCmd).Standalone()

	authCmd.AddCommand(auth_gitCredentialCmd)
}
