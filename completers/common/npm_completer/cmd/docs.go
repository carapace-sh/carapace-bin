package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/npm_completer/cmd/action"
	"github.com/spf13/cobra"
)

var docsCmd = &cobra.Command{
	Use:   "docs",
	Short: "Open documentation for a package in a web browser",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docsCmd).Standalone()
	docsCmd.Flags().String("browser", "", "browser to use")
	addWorkspaceFlags(docsCmd)

	rootCmd.AddCommand(docsCmd)

	carapace.Gen(docsCmd).PositionalAnyCompletion(
		action.ActionPackageNames(docsCmd),
	)
}
