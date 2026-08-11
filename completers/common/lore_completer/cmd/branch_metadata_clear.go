package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var branch_metadata_clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear metadata from the branch",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_metadata_clearCmd).Standalone()

	branch_metadata_clearCmd.Flags().String("branch", "", "Branch name (uses current branch if not specified)")
	branch_metadata_clearCmd.Flags().BoolP("help", "h", false, "Print help")
	branch_metadataCmd.AddCommand(branch_metadata_clearCmd)
}
