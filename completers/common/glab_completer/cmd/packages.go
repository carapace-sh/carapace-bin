package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var packagesCmd = &cobra.Command{
	Use:   "packages <command> [flags]",
	Short: "Manage packages in the GitLab package registry.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(packagesCmd).Standalone()

	packagesCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository. Can use either `OWNER/REPO` or `GROUP/NAMESPACE/REPO` format. Also accepts full URL or Git URL.")
	rootCmd.AddCommand(packagesCmd)

	carapace.Gen(packagesCmd).FlagCompletion(carapace.ActionMap{
		"repo": action.ActionRepo(packagesCmd),
	})
}
