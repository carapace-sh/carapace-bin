package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_linkCmd = &cobra.Command{
	Use:   "link",
	Short: "Links a local package as a dependency",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_linkCmd).Standalone()

	helpCmd.AddCommand(help_linkCmd)
}
