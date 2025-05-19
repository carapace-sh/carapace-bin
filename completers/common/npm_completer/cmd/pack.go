package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/npm_completer/cmd/action"
	"github.com/spf13/cobra"
)

var packCmd = &cobra.Command{
	Use:   "pack",
	Short: "Create a tarball from a package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(packCmd).Standalone()
	packCmd.Flags().Bool("dry-run", false, "only report changes")
	packCmd.Flags().Bool("json", false, "output as json")
	addWorkspaceFlags(packCmd)

	rootCmd.AddCommand(packCmd)

	carapace.Gen(packCmd).PositionalAnyCompletion(
		action.ActionPackages(packCmd),
	)
}
