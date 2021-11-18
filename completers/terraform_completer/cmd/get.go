package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Install or upgrade remote Terraform modules",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(getCmd).Standalone()

	getCmd.Flags().Bool("no-color", false, "Disable text coloring in the output.")
	getCmd.Flags().Bool("update", false, "Check already-downloaded modules for available updates")
	rootCmd.AddCommand(getCmd)

	carapace.Gen(getCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
