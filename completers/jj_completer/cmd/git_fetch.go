package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var git_fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetch from a Git remote",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(git_fetchCmd).Standalone()

	git_fetchCmd.Flags().Bool("all-remotes", false, "Fetch from all remotes")
	git_fetchCmd.Flags().StringSlice("branch", []string{}, "Fetch only some of the branches")
	git_fetchCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	git_fetchCmd.Flags().StringSlice("remote", []string{}, "The remote to fetch from (only named remotes are supported, can be repeated)")
	gitCmd.AddCommand(git_fetchCmd)
}
