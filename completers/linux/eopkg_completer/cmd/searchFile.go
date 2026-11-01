package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var searchFileCmd = &cobra.Command{
	Use:     "search-file <path>",
	Aliases: []string{"sf"},
	Short:   "locate the package which is considered to be the owner of the specified path on disk",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(searchFileCmd).Standalone()

	searchFileCmd.Flags().BoolP("long", "l", false, "show detailed information about matching packages")
	searchFileCmd.Flags().BoolP("quiet", "q", false, "terse output only showing the package name")

	rootCmd.AddCommand(searchFileCmd)

	carapace.Gen(searchFileCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
