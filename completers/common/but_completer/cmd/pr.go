package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var prCmd = &cobra.Command{
	Use:     "pr",
	Short:   "Commands for creating and managing pull requests on a forge",
	Aliases: []string{"review", "mr"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(prCmd).Standalone()

	prCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(prCmd)
}
