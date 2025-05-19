package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var sshKeyCmd = &cobra.Command{
	Use:   "ssh-key <command>",
	Short: "Manage SSH keys registered with your GitLab account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sshKeyCmd).Standalone()

	sshKeyCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository. Can use either `OWNER/REPO` or `GROUP/NAMESPACE/REPO` format. Also accepts full URL or Git URL.")
	rootCmd.AddCommand(sshKeyCmd)

	carapace.Gen(sshKeyCmd).FlagCompletion(carapace.ActionMap{
		"repo": action.ActionRepo(variableCmd),
	})
}
