package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var listSourcesCmd = &cobra.Command{
	Use:     "list-sources",
	Aliases: []string{"ls"},
	Short:   "output all source packages available for emerge operations",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listSourcesCmd).Standalone()

	listSourcesCmd.Flags().BoolP("long", "l", false, "show detailed information on each source package")

	rootCmd.AddCommand(listSourcesCmd)
}
