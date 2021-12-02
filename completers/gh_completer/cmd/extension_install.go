package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var extension_installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install a gh extension from a repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(extension_installCmd).Standalone()
	extensionCmd.AddCommand(extension_installCmd)

	carapace.Gen(extension_installCmd).PositionalCompletion(
		action.ActionOwnerRepositories(extension_installCmd),
	)
}
