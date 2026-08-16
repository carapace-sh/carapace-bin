package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var git_cloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "Clone a Git repository.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(git_cloneCmd).Standalone()

	gitCmd.AddCommand(git_cloneCmd)
}
