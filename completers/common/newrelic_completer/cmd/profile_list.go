package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var profile_listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List the profiles available",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(profile_listCmd).Standalone()
	profile_listCmd.Flags().BoolP("show-keys", "s", false, "list the profiles on your keychain")
	profileCmd.AddCommand(profile_listCmd)
}
