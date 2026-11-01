package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "stack <command>",
	Short: "Manage stacked branches and pull requests",
	Long:  "https://github.com/github/gh-stack",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()
	rootCmd.AddGroup(
		&cobra.Group{ID: "stack", Title: "stack commands"},
		&cobra.Group{ID: "nav", Title: "navigation commands"},
		&cobra.Group{ID: "remote", Title: "remote commands"},
		&cobra.Group{ID: "utils", Title: "utility commands"},
	)

}
