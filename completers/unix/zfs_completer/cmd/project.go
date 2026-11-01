package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var projectCmd = &cobra.Command{
	Use:     "project [-c|-C|-l|-p id|-s] [-r] [-d] [-0] [-k] file|directory...",
	Short:   "manage project IDs and inheritance",
	GroupID: "misc",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(projectCmd).Standalone()

	projectCmd.Flags().BoolS("0", "0", false, "delimit with NUL instead of newline")
	projectCmd.Flags().BoolS("C", "C", false, "clear project inherit flag")
	projectCmd.Flags().BoolS("c", "c", false, "check project ID and flags")
	projectCmd.Flags().BoolS("d", "d", false, "show/check the directory itself, not its children")
	projectCmd.Flags().BoolS("k", "k", false, "keep the project ID unchanged when clearing")
	projectCmd.Flags().BoolS("l", "l", false, "list project ID and inherit flag")
	projectCmd.Flags().StringS("p", "p", "", "set project ID")
	projectCmd.Flags().BoolS("r", "r", false, "apply recursively")
	projectCmd.Flags().BoolS("s", "s", false, "set project inherit flag")

	rootCmd.AddCommand(projectCmd)

	carapace.Gen(projectCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
