package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/npm"
	"github.com/spf13/cobra"
)

var bugsCmd = &cobra.Command{
	Use:     "bugs",
	Short:   "Opens the bug tracker URL of a package in the default browser",
	Aliases: []string{"issues"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bugsCmd).Standalone()

	bugsCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	bugsCmd.Flags().String("registry", "", "")
	rootCmd.AddCommand(bugsCmd)

	carapace.Gen(bugsCmd).PositionalCompletion(
		npm.ActionPackageSearch(""),
	)
}
