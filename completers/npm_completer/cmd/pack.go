package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/npm_completer/cmd/action"
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
	packCmd.Flags().StringP("workspace", "w", "", "Enable running a command in the context of the given workspace")
	packCmd.Flags().Bool("workspaces", false, "Enable running a command in the context fo all workspaces")

	rootCmd.AddCommand(packCmd)

	carapace.Gen(packCmd).PositionalAnyCompletion(
		action.ActionPackages(packCmd),
	)
}
