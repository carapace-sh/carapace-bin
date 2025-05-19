package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var submoduleCmd = &cobra.Command{
	Use:     "submodule",
	Short:   "Only print error messages",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_main].ID,
}

func init() {
	carapace.Gen(submoduleCmd).Standalone()
	submoduleCmd.PersistentFlags().Bool("quiet", false, "Only print error messages")

	rootCmd.AddCommand(submoduleCmd)
}
