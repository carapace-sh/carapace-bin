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
	submodule_updateCmd.Flags().Bool("checkout", false, "checkout commit of superproject as detached HEAD")
	submodule_updateCmd.Flags().String("depth", "", "create a shallow clone")
	submodule_updateCmd.Flags().Bool("dissociate", false, "use --reference only while cloning")
	submodule_updateCmd.Flags().StringArray("filter", nil, "object filtering")
	submodule_updateCmd.Flags().BoolP("force", "f", false, "throw away local changes")
	submodule_updateCmd.Flags().BoolP("init", "i", false, "initialize all submodules")
	submodule_updateCmd.Flags().StringP("jobs", "j", "", "clone new submodules in parallel")
	submodule_updateCmd.Flags().BoolP("merge", "m", false, "merge commit of superproject into submodule")
	submodule_updateCmd.Flags().BoolP("no-fetch", "N", false, "don't fetch new objects from the remote")
	submodule_updateCmd.Flags().Bool("no-recommend-shallow", false, "ignore configured shallow recommendation")
	submodule_updateCmd.Flags().Bool("progress", false, "force cloning progress")
	submodule_updateCmd.Flags().BoolP("rebase", "r", false, "rebase current branch onto commit of superproject")
	submodule_updateCmd.Flags().Bool("recommend-shallow", false, "initial clone will use recommended configuration")
	submodule_updateCmd.Flags().Bool("recursive", false, "traverse submodules recursively")
	submodule_updateCmd.Flags().String("ref-format", "", "specify the reference storage format to use")
	submodule_updateCmd.Flags().String("reference", "", "reference repository")
	submodule_updateCmd.Flags().Bool("remote", false, "use submodules's remote tracking branch to update")
	submodule_updateCmd.Flags().Bool("require-init", false, "disallow cloning into non-empty directory, implies --init")
	submodule_updateCmd.Flags().Bool("single-branch", false, "clone only one branch, HEAD or --branch")

	submoduleCmd.AddCommand(submodule_updateCmd)

	carapace.Gen(submodule_updateCmd).FlagCompletion(carapace.ActionMap{
		"filter":     git.ActionObjectFilters(),
		"ref-format": carapace.ActionValuesDescribed("files", "loose files with packed-refs", "reftable", "reftable format"),
		"reference":  carapace.ActionDirectories(),
	})

	carapace.Gen(submodule_updateCmd).PositionalAnyCompletion(
		git.ActionSubmodulePaths().FilterArgs(),
	)

}
