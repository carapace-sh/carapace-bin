package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var allCmd = &cobra.Command{
	Use:     "all",
	Aliases: []string{"build"},
	Short:   "Build the project",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(allCmd).Standalone()
	rootCmd.AddCommand(allCmd)
}
