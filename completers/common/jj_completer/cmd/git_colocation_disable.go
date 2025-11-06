package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var git_colocation_disableCmd = &cobra.Command{
	Use:   "disable",
	Short: "Convert into a non-colocated Jujutsu/Git repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(git_colocation_disableCmd).Standalone()

	git_colocation_disableCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	git_colocationCmd.AddCommand(git_colocation_disableCmd)
}
