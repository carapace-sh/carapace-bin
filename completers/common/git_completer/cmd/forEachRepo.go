package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var forEachRepoCmd = &cobra.Command{
	Use:     "for-each-repo",
	Short:   "Run a Git command on a list of repositories",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_low_level_interrogator].ID,
}

func init() {
	carapace.Gen(forEachRepoCmd).Standalone()
	forEachRepoCmd.Flags().SetInterspersed(false)

	forEachRepoCmd.Flags().String("config", "", "use the given config variable as a multi-valued list storing absolute path names.")
	forEachRepoCmd.Flags().Bool("keep-going", false, "continue with the remaining repositories if the command failed on a repository")
	rootCmd.AddCommand(forEachRepoCmd)

	// TODO complete config

	carapace.Gen(forEachRepoCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin("git"),
	)

	carapace.Gen(forEachRefCmd).DashAnyCompletion(
		carapace.ActionPositional(forEachRepoCmd),
	)
}
