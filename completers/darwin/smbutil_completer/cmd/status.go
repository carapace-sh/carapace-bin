package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Resolve hostname to NetBIOS workgroup and system name",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(statusCmd).Standalone()
	rootCmd.AddCommand(statusCmd)

	statusCmd.Flags().BoolS("a", "a", false, "Display all NetBIOS names")
	statusCmd.Flags().BoolS("e", "e", false, "Percent escape NetBIOS names")

	carapace.Gen(statusCmd).PositionalCompletion(
		carapace.ActionValues(),
	)
}
