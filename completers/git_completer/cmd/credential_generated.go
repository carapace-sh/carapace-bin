package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var credentialCmd = &cobra.Command{
	Use:     "credential",
	Short:   "Retrieve and store user credentials",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_low_level_helper].ID,
}

func init() {
	carapace.Gen(credentialCmd).Standalone()

	rootCmd.AddCommand(credentialCmd)
}
