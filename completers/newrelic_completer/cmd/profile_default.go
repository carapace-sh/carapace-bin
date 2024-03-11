package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var profile_defaultCmd = &cobra.Command{
	Use:   "default",
	Short: "Set the default profile name",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(profile_defaultCmd).Standalone()
	profileCmd.AddCommand(profile_defaultCmd)
}
