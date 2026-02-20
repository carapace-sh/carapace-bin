package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var listSourcesCmd = &cobra.Command{
	Use:     "list-sources",
	Short:   "List available sources",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listSourcesCmd).Standalone()

	listSourcesCmd.Flags().BoolP("long", "l", false, "Show in long format")

	rootCmd.AddCommand(listSourcesCmd)
}
