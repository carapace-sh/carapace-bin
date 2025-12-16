package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pnpm",
	Short: "Fast, disk space efficient package manager",
	Long:  "https://pnpm.io/",
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
		&cobra.Group{ID: "manage", Title: "Manage Commands"},
		&cobra.Group{ID: "review", Title: "Review Commands"},
		&cobra.Group{ID: "run", Title: "Run Commands"},
		&cobra.Group{ID: "store", Title: "Store Commands"},
	)

	rootCmd.Flags().BoolP("help", "h", false, "show help")
	rootCmd.Flags().BoolP("recursive", "r", false, "Run the command for each project in the workspace")
	rootCmd.Flags().BoolP("version", "v", false, "show version")
}
