package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var profile_deleteCmd = &cobra.Command{
	Use:     "delete",
	Short:   "Delete a profile",
	Aliases: []string{"remove", "rm"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(profile_deleteCmd).Standalone()
	profileCmd.AddCommand(profile_deleteCmd)
}
