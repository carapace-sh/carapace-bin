package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var submodule_updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update the registered submodule",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(submodule_updateCmd).Standalone()
	// TODO find flag descriptions
	submodule_updateCmd.Flags().Bool("checkout", false, "checkout commit of superproject as detached HEAD")
	submodule_updateCmd.Flags().String("depth", "", "create a shallow clone")
	submodule_updateCmd.Flags().BoolP("force", "f", false, "throw away local changes")
	submodule_updateCmd.Flags().Bool("init", false, "initialize all submodules")
	submodule_updateCmd.Flags().String("jobs", "", "clone new submodules in parallel")
	submodule_updateCmd.Flags().Bool("merge", false, "merge commit of superproject into submodule")
	submodule_updateCmd.Flags().BoolP("no-Fetch", "N", false, "don't fetch from remote")
	submodule_updateCmd.Flags().Bool("no-recommend-shallow", false, "ignore configured shallow recommendation")
	submodule_updateCmd.Flags().Bool("rebase", false, "rebase current branch onto commit of superproject")
	submodule_updateCmd.Flags().Bool("recommend-shallow", false, "initial clone will use recommended configuration")
	submodule_updateCmd.Flags().Bool("recursive", false, "traverse submodules recursively")
	submodule_updateCmd.Flags().String("reference", "", "reference repository")
	submodule_updateCmd.Flags().Bool("remote", false, "use submodules's remote tracking branch to update")

	submoduleCmd.AddCommand(submodule_updateCmd)

	carapace.Gen(submodule_updateCmd).FlagCompletion(carapace.ActionMap{
		"reference": carapace.ActionDirectories(),
	})

	carapace.Gen(submodule_updateCmd).PositionalAnyCompletion(
		git.ActionSubmodulePaths().FilterArgs(),
	)

}
