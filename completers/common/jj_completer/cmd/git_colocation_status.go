package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var git_colocation_statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show the current colocation status",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(git_colocation_statusCmd).Standalone()

	git_colocation_statusCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	git_colocationCmd.AddCommand(git_colocation_statusCmd)
}
