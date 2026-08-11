package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var link_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove the link at the given point in the repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(link_removeCmd).Standalone()

	link_removeCmd.Flags().BoolP("help", "h", false, "Print help")
	linkCmd.AddCommand(link_removeCmd)
}
