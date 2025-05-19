package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "yarn",
	Short: "Yarn is a package manager that doubles down as project manager",
	Long:  "https://yarnpkg.com/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.AddGroup(
		&cobra.Group{ID: "general", Title: "General commands"},
		&cobra.Group{ID: "npm-related", Title: "Npm-related commands"},
		&cobra.Group{ID: "plugin-related", Title: "Plugin-related commands"},
		&cobra.Group{ID: "workspace-related", Title: "Workspace-related commands"},
	)

	rootCmd.PersistentFlags().BoolP("help", "h", false, "show help")

}
