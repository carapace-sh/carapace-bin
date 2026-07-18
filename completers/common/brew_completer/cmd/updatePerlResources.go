package cmd

import (
	"github.com/carapace-sh/carapace"
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

	updatePerlResourcesCmd.Flags().Bool("debug", false, "Display any debugging information.")
	updatePerlResourcesCmd.Flags().Bool("help", false, "Show this message.")
	updatePerlResourcesCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	updatePerlResourcesCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	rootCmd.AddCommand(updatePerlResourcesCmd)
}
