package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var groupRemoveCmd = &cobra.Command{
	Use:   "remove [options] <group-spec>...",
	Short: "remove groups",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(groupRemoveCmd).Standalone()

	groupRemoveCmd.Flags().Bool("no-packages", false, "operate on groups without manipulating packages")
	groupRemoveCmd.Flags().Bool("with-optional", false, "include optional packages")

	groupCmd.AddCommand(groupRemoveCmd)
}
