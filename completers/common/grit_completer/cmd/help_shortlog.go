package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_shortlogCmd = &cobra.Command{
	Use:   "shortlog",
	Short: "List the commits on this branch that aren't on the target branch yet",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_shortlogCmd).Standalone()

	helpCmd.AddCommand(help_shortlogCmd)
}
