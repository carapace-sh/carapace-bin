package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var gpgKeyCmd = &cobra.Command{
	Use:   "gpg-key <command>",
	Short: "Manage GPG keys registered with your GitLab account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gpgKeyCmd).Standalone()

	gpgKeyCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository. Can use either `OWNER/REPO` or `GROUP/NAMESPACE/REPO` format. Also accepts full URL or Git URL.")
	rootCmd.AddCommand(gpgKeyCmd)

	carapace.Gen(gpgKeyCmd).FlagCompletion(carapace.ActionMap{
		"repo": action.ActionRepo(gpgKeyCmd),
	})
}
