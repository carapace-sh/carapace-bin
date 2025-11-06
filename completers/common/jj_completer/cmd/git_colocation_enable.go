package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var git_colocation_enableCmd = &cobra.Command{
	Use:   "enable",
	Short: "Convert into a colocated Jujutsu/Git repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(git_colocation_enableCmd).Standalone()

	git_colocation_enableCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	git_colocationCmd.AddCommand(git_colocation_enableCmd)
}
