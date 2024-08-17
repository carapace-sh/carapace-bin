package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var refsCmd = &cobra.Command{
	Use:     "refs",
	Short:   "Low-level access to refs",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_manipulator].ID,
}

func init() {
	carapace.Gen(refsCmd).Standalone()

	rootCmd.AddCommand(refsCmd)
}
