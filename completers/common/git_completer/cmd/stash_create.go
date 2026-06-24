package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var stash_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a stash entry without reflog entry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stash_createCmd).Standalone()

	stashCmd.AddCommand(stash_createCmd)
}
