package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/brew"
	"github.com/spf13/cobra"
)

var homeCmd = &cobra.Command{
	Use:     "home",
	Short:   "Open a <formula> or <cask>'s homepage in a browser, or open Homebrew's own homepage if no argument is provided",
	GroupID: "main",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(homeCmd).Standalone()

	homeCmd.Flags().Bool("cask", false, "Treat all named arguments as casks.")
	homeCmd.Flags().BoolP("debug", "d", false, "Display any debugging information.")
	homeCmd.Flags().Bool("formula", false, "Treat all named arguments as formulae.")
	homeCmd.Flags().BoolP("help", "h", false, "Show this message.")
	homeCmd.Flags().BoolP("quiet", "q", false, "Make some output more quiet.")
	homeCmd.Flags().BoolP("verbose", "v", false, "Make some output more verbose.")
	rootCmd.AddCommand(homeCmd)

	carapace.Gen(homeCmd).PositionalAnyCompletion(
		carapace.Batch(
			brew.ActionAllCasks().Unless(homeCmd.Flag("formula").Changed),
			brew.ActionAllFormulae().Unless(homeCmd.Flag("cask").Changed),
		).ToA().FilterArgs(),
	)
}
