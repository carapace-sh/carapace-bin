package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var linkCmd = &cobra.Command{
	Use:     "link",
	Short:   "Links a local package as a dependency",
	Aliases: []string{"ln"},
	GroupID: "manage",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(linkCmd).Standalone()

	linkCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(linkCmd)

	carapace.Gen(linkCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
