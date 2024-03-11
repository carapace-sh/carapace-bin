package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var citoolCmd = &cobra.Command{
	Use:     "citool",
	Short:   "Graphical alternative to git-commit",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_main].ID,
}

func init() {
	carapace.Gen(citoolCmd).Standalone()
	rootCmd.AddCommand(citoolCmd)
}
