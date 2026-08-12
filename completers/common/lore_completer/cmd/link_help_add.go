package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var link_help_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Link to the given point in the repository and subpath from the given repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(link_help_addCmd).Standalone()

	link_helpCmd.AddCommand(link_help_addCmd)
}
