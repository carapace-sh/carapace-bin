package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/npm_completer/cmd/action"
	"github.com/spf13/cobra"
)

var repoCmd = &cobra.Command{
	Use:   "repo",
	Short: "Open package repository page in the browser",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repoCmd).Standalone()
	repoCmd.Flags().String("browser", "", "browser to user")
	repoCmd.Flags().Bool("no-browser", false, "print urls to terminal instead of opening browser")
	repoCmd.Flags().String("registry", "", "base URL of the npm registry")
	addWorkspaceFlags(repoCmd)

	rootCmd.AddCommand(repoCmd)

	carapace.Gen(repoCmd).FlagCompletion(carapace.ActionMap{
		"browser": carapace.ActionFiles(),
	})

	carapace.Gen(repoCmd).PositionalAnyCompletion(
		action.ActionPackageNames(repoCmd),
	)
}
