package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_starsCmd = &cobra.Command{
	Use:   "stars",
	Short: "Lists all packages starred by a specific user",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_starsCmd).Standalone()

	helpCmd.AddCommand(help_starsCmd)
}
