package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/brew"
	"github.com/spf13/cobra"
)

var updatePerlResourcesCmd = &cobra.Command{
	Use:     "update-perl-resources",
	Short:   "Update versions for CPAN resource blocks in <formula>",
	GroupID: "developer",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(updatePerlResourcesCmd).Standalone()

	updatePerlResourcesCmd.Flags().BoolP("debug", "d", false, "Display any debugging information.")
	updatePerlResourcesCmd.Flags().BoolP("help", "h", false, "Show this message.")
	updatePerlResourcesCmd.Flags().BoolP("quiet", "q", false, "Make some output more quiet.")
	updatePerlResourcesCmd.Flags().BoolP("verbose", "v", false, "Make some output more verbose.")
	rootCmd.AddCommand(updatePerlResourcesCmd)

	carapace.Gen(updatePerlResourcesCmd).PositionalAnyCompletion(
		brew.ActionAllFormulae(),
	)
}
