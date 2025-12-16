package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tea",
	Short: "command line tool to interact with Gitea",
	Long:  "https://gitea.com/gitea/tea",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute(opts ...func(cmd *cobra.Command)) error {
	for _, opt := range opts {
		opt(rootCmd)
	}
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()
	rootCmd.AddGroup(
		&cobra.Group{ID: "SETUP", Title: ""},
		&cobra.Group{ID: "MISCELLANEOUS", Title: ""},
		&cobra.Group{ID: "ENTITIES", Title: ""},
		&cobra.Group{ID: "HELPERS", Title: ""},
	)
}
