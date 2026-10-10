package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var securefileCmd = &cobra.Command{
	Use:   "securefile <command> [flags]",
	Short: "Manage secure files for a project.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securefileCmd).Standalone()

	securefileCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository. You can use either OWNER/REPO or GROUP/NAMESPACE/REPO. The full URL or Git URL is also accepted.")
	rootCmd.AddCommand(securefileCmd)

	carapace.Gen(securefileCmd).FlagCompletion(carapace.ActionMap{
		"repo": action.ActionRepo(securefileCmd),
	})
}
