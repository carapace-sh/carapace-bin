package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var shortlogCmd = &cobra.Command{
	Use:     "shortlog",
	Short:   "List the commits on this branch that aren't on the target branch yet",
	Aliases: []string{"sl"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(shortlogCmd).Standalone()

	shortlogCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(shortlogCmd)
}
