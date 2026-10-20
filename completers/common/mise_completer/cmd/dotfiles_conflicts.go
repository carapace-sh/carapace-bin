package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var dotfiles_conflictsCmd = &cobra.Command{
	Use:   "conflicts",
	Short: "Compare both sides of a sync conflict",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dotfiles_conflictsCmd).Standalone()

	dotfiles_conflictsCmd.Flags().Bool("difftool", false, "Open the comparison in Git's configured diff or merge tool")
	dotfiles_conflictsCmd.Flags().BoolP("help", "h", false, "Print help")
	dotfiles_conflictsCmd.Flags().String("tool", "", "Use this Git diff tool instead of the configured default")
	dotfilesCmd.AddCommand(dotfiles_conflictsCmd)

	carapace.Gen(dotfiles_conflictsCmd).FlagCompletion(carapace.ActionMap{
		"tool": git.ActionDiffTools(),
	})
}
