package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var sh_setupCmd = &cobra.Command{
	Use:     "sh-setup",
	Short:   "Common Git shell script setup code",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_low_level_helper].ID,
}

func init() {
	carapace.Gen(sh_setupCmd).Standalone()

	rootCmd.AddCommand(sh_setupCmd)
}
