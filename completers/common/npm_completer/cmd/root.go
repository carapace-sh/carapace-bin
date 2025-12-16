package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/npm"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "npm",
	Short: "the package manager for JavaScript",
	Long:  "https://www.npmjs.com/",
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
	rootCmd.PersistentFlags().String("registry", "", "set registry to use")
}

func addWorkspaceFlags(cmd *cobra.Command) {
	cmd.Flags().StringArrayP("workspace", "w", []string{""}, "Enable running a command in the context of the given workspace")
	cmd.Flags().Bool("workspaces", false, "Enable running a command in the context fo all workspaces")

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"workspace": npm.ActionWorkspaces(),
	})
}
