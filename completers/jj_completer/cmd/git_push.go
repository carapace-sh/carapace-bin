package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var git_pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Push to a Git remote",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(git_pushCmd).Standalone()

	git_pushCmd.Flags().Bool("all", false, "Push all branches (including deleted branches)")
	git_pushCmd.Flags().StringSliceP("branch", "b", []string{}, "Push only this branch (can be repeated)")
	git_pushCmd.Flags().StringSliceP("change", "c", []string{}, "Push this commit by creating a branch based on its change ID (can be repeated)")
	git_pushCmd.Flags().Bool("deleted", false, "Push all deleted branches")
	git_pushCmd.Flags().Bool("dry-run", false, "Only display what will change on the remote")
	git_pushCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	git_pushCmd.Flags().String("remote", "", "The remote to push to (only named remotes are supported)")
	git_pushCmd.Flags().StringSliceP("revisions", "r", []string{}, "Push branches pointing to these commits")
	gitCmd.AddCommand(git_pushCmd)

	carapace.Gen(git_pushCmd).FlagCompletion(carapace.ActionMap{
		"branch":    jj.ActionLocalBranches(),
		"change":    carapace.ActionValues(), // TODO
		"remote":    jj.ActionRemotes(),
		"revisions": jj.ActionRevSets(jj.RevOption{}.Default()),
	})
}
