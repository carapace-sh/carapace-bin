package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zfs"
	"github.com/spf13/cobra"
)

var unjailCmd = &cobra.Command{
	Use:     "unjail jailid filesystem",
	Short:   "detach a filesystem from a jail",
	GroupID: "misc",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(unjailCmd).Standalone()

	rootCmd.AddCommand(unjailCmd)

	carapace.Gen(unjailCmd).PositionalCompletion(
		carapace.ActionValues(),
		zfs.ActionFilesystems(),
	)
}
