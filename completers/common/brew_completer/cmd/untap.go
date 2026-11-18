package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/brew"
	"github.com/spf13/cobra"
)

var untapCmd = &cobra.Command{
	Use:     "untap",
	Short:   "Remove a tapped formula repository",
	GroupID: "main",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(untapCmd).Standalone()

	untapCmd.Flags().BoolP("debug", "d", false, "Display any debugging information.")
	untapCmd.Flags().BoolP("force", "f", false, "Untap even if formulae or casks from this tap are currently installed.")
	untapCmd.Flags().BoolP("help", "h", false, "Show this message.")
	untapCmd.Flags().BoolP("quiet", "q", false, "Make some output more quiet.")
	untapCmd.Flags().BoolP("verbose", "v", false, "Make some output more verbose.")
	rootCmd.AddCommand(untapCmd)

	carapace.Gen(untapCmd).PositionalAnyCompletion(
		brew.ActionInstalledTaps().FilterArgs(),
	)
}
