package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backendsCmd = &cobra.Command{
	Use:     "backends",
	Short:   "Manage backends",
	Aliases: []string{"b", "backend", "backend-list"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backendsCmd).Standalone()

	backendsCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(backendsCmd)
}
