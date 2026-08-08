package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var skills_installCmd = &cobra.Command{
	Use:   "install [name]",
	Short: "Install glab's bundled agent skills. (EXPERIMENTAL)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(skills_installCmd).Standalone()

	skills_installCmd.Flags().BoolP("force", "f", false, "Overwrite existing skill files.")
	skills_installCmd.Flags().BoolP("global", "g", false, "Install skills at user scope (~/.agents/skills/).")
	skills_installCmd.Flags().String("path", "", "Install skills to the directory at <path>.")
	skillsCmd.AddCommand(skills_installCmd)

	carapace.Gen(skills_installCmd).FlagCompletion(carapace.ActionMap{
		"path": carapace.ActionDirectories(),
	})
}
