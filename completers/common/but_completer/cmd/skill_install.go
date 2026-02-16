package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/traverse"
	"github.com/spf13/cobra"
)

var skill_installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install the GitButler CLI skill files for Coding agents",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(skill_installCmd).Standalone()

	skill_installCmd.Flags().BoolP("detect", "d", false, "Automatically detect where to install by finding existing installation")
	skill_installCmd.Flags().BoolP("global", "g", false, "Install the skill globally instead of in the current repository")
	skill_installCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	skill_installCmd.Flags().StringP("path", "p", "", "Custom path where to install the skill (relative to repository root or absolute). Outside a repository, relative paths require --global")
	skillCmd.AddCommand(skill_installCmd)

	carapace.Gen(skill_installCmd).FlagCompletion(carapace.ActionMap{
		"path": carapace.ActionDirectories(),
	})

	carapace.Gen(skill_installCmd).PositionalCompletion(
		carapace.ActionFiles().ChdirF(traverse.GitWorkTree),
	)
}
