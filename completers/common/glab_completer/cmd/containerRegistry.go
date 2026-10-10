package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var containerRegistryCmd = &cobra.Command{
	Use:     "container-registry <command> [flags]",
	Short:   "Work with GitLab container registries.",
	Aliases: []string{"cr"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(containerRegistryCmd).Standalone()

	containerRegistryCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository. You can use either OWNER/REPO or GROUP/NAMESPACE/REPO. The full URL or Git URL is also accepted.")
	rootCmd.AddCommand(containerRegistryCmd)

	carapace.Gen(containerRegistryCmd).FlagCompletion(carapace.ActionMap{
		"repo": action.ActionRepo(containerRegistryCmd),
	})
}
