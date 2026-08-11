package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var link_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Link to the given point in the repository and subpath from the given repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(link_addCmd).Standalone()

	link_addCmd.Flags().Bool("disable-branching", false, "Disable automatic branch creation in the linked repository")
	link_addCmd.Flags().BoolP("help", "h", false, "Print help")
	link_addCmd.Flags().String("pin", "", "Branch or specific revision to pin the link to, defaulting to latest on the main branch")
	linkCmd.AddCommand(link_addCmd)
}
