package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pnpm"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pnpm",
	Short: "Fast, disk space efficient package manager",
	Long:  "https://pnpm.io/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.AddGroup(
		&cobra.Group{ID: "manage", Title: "Manage Commands"},
		&cobra.Group{ID: "review", Title: "Review Commands"},
		&cobra.Group{ID: "run", Title: "Run Commands"},
		&cobra.Group{ID: "store", Title: "Store Commands"},
		&cobra.Group{ID: "other", Title: "Other Commands"},
	)

	rootCmd.Flags().Bool("color", false, "Controls colors in the output")
	rootCmd.Flags().String("filter", "", "set filter")
	rootCmd.Flags().String("filter-prod", "", "Restricts the scope to package names matching the given pattern")
	rootCmd.Flags().BoolP("help", "h", false, "show help")
	rootCmd.Flags().String("loglevel", "", "What level of logs to report")
	rootCmd.Flags().Bool("no-color", false, "Controls colors in the output")
	rootCmd.Flags().BoolP("recursive", "r", false, "Run the command for each project in the workspace")
	rootCmd.Flags().BoolP("version", "v", false, "show version")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"filter":      pnpm.ActionFilter(),
		"filter-prod": pnpm.ActionFilter(),
		"loglevel":    pnpm.ActionLoglevel(),
	})
}

func addWorkspaceFlags(cmd *cobra.Command) {
	cmd.Flags().StringArrayP("workspace", "w", []string{""}, "Enable running a command in the context of the given workspace")
	cmd.Flags().Bool("workspaces", false, "Enable running a command in the context of all workspaces")

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"workspace": carapace.Batch(
			pnpm.ActionWorkspaces(),
			pnpm.ActionWorkspaceDependencies(),
		).ToA(),
	})
}
