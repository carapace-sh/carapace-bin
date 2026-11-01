package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gpgKey_addCmd = &cobra.Command{
	Use:   "add [key-file]",
	Short: "Add a GPG key to your GitLab account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gpgKey_addCmd).Standalone()

	gpgKeyCmd.AddCommand(gpgKey_addCmd)
}
