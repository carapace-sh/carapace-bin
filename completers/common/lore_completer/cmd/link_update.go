package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var link_updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update the link to a new pin",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(link_updateCmd).Standalone()

	link_updateCmd.Flags().BoolP("help", "h", false, "Print help")
	link_updateCmd.Flags().String("pin", "", "Branch or specific revision to pin the link to, defaulting to latest on the current branch")
	linkCmd.AddCommand(link_updateCmd)
}
