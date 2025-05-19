package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var clusterCmd = &cobra.Command{
	Use:   "cluster <command> [flags]",
	Short: "Manage GitLab Agents for Kubernetes and their clusters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clusterCmd).Standalone()

	clusterCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository. Can use either `OWNER/REPO` or `GROUP/NAMESPACE/REPO` format. Also accepts full URL or Git URL.")
	rootCmd.AddCommand(clusterCmd)

	carapace.Gen(clusterCmd).FlagCompletion(carapace.ActionMap{
		"repo": action.ActionRepo(clusterCmd),
	})
}
