package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var help_lintersCmd = &cobra.Command{
	Use:   "linters",
	Short: "Help about linters",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_lintersCmd).Standalone()

	helpCmd.AddCommand(help_lintersCmd)
}
