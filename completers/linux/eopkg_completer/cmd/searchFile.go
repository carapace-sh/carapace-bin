package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var searchFileCmd = &cobra.Command{
	Use:     "search-file",
	Short:   "Search for a file in installed packages",
	Aliases: []string{"sf"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(searchFileCmd).Standalone()

	searchFileCmd.Flags().BoolP("long", "l", false, "Show in long format")
	searchFileCmd.Flags().BoolP("quiet", "q", false, "Show only package name")

	rootCmd.AddCommand(searchFileCmd)
}
