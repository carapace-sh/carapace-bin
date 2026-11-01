package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var securityCmd = &cobra.Command{
	Use:   "security <command> [flags]",
	Short: "Manage GitLab security scan profiles for a project. (EXPERIMENTAL)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityCmd).Standalone()

	securityCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository. You can use either OWNER/REPO or GROUP/NAMESPACE/REPO. The full URL or Git URL is also accepted.")
	rootCmd.AddCommand(securityCmd)

	carapace.Gen(securityCmd).FlagCompletion(carapace.ActionMap{
		"repo": action.ActionRepo(securityCmd),
	})
}
