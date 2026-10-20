package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var skills_syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Link the active tools' skills where an agent looks for them",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(skills_syncCmd).Standalone()

	skills_syncCmd.Flags().String("dir", "", "The directory to link skills into")
	skills_syncCmd.Flags().BoolP("global", "g", false, "Link into the skills.dir directory under your home directory instead of the project's")
	skills_syncCmd.Flags().BoolP("help", "h", false, "Print help")
	skills_syncCmd.Flags().Bool("prune", false, "Remove links mise made for skills that are no longer active")
	skillsCmd.AddCommand(skills_syncCmd)

	carapace.Gen(skills_syncCmd).FlagCompletion(carapace.ActionMap{
		"dir": carapace.ActionDirectories(),
	})
}
