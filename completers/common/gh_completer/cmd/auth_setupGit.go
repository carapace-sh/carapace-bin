package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var auth_setupGitCmd = &cobra.Command{
	Use:   "setup-git",
	Short: "Setup git with GitHub CLI",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auth_setupGitCmd).Standalone()

	auth_setupGitCmd.Flags().BoolP("force", "f", false, "Force setup even if the host is not known. Must be used in conjunction with `--hostname`")
	auth_setupGitCmd.Flags().StringP("hostname", "h", "", "The hostname to configure git for")
	authCmd.AddCommand(auth_setupGitCmd)

	carapace.Gen(auth_setupGitCmd).FlagCompletion(carapace.ActionMap{
		"hostname": action.ActionConfigHosts(),
	})
}
