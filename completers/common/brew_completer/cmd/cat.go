package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/brew"
	"github.com/spf13/cobra"
)

var catCmd = &cobra.Command{
	Use:     "cat",
	Short:   "Display the source of a <formula> or <cask>",
	GroupID: "developer",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(catCmd).Standalone()

	catCmd.Flags().Bool("cask", false, "Treat all named arguments as casks.")
	catCmd.Flags().BoolP("debug", "d", false, "Display any debugging information.")
	catCmd.Flags().Bool("formula", false, "Treat all named arguments as formulae.")
	catCmd.Flags().BoolP("help", "h", false, "Show this message.")
	catCmd.Flags().BoolP("quiet", "q", false, "Make some output more quiet.")
	catCmd.Flags().BoolP("verbose", "v", false, "Make some output more verbose.")
	rootCmd.AddCommand(catCmd)

	carapace.Gen(catCmd).PositionalAnyCompletion(
		carapace.Batch(
			brew.ActionAllCasks().Unless(catCmd.Flag("formula").Changed),
			brew.ActionAllFormulae().Unless(catCmd.Flag("cask").Changed),
		).ToA().FilterArgs(),
	)
}
