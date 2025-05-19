package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var profileCmd = &cobra.Command{
	Use:     "profile",
	Short:   "Manage the authentication profiles for this tool",
	Aliases: []string{"profiles"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(profileCmd).Standalone()
	rootCmd.AddCommand(profileCmd)
}
