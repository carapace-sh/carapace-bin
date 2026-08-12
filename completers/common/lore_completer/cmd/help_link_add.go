package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_link_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Link to the given point in the repository and subpath from the given repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_link_addCmd).Standalone()

	help_linkCmd.AddCommand(help_link_addCmd)
}
