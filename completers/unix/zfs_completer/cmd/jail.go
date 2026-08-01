package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zfs"
	"github.com/spf13/cobra"
)

var jailCmd = &cobra.Command{
	Use:     "jail jailid filesystem",
	Short:   "attach a filesystem to a jail",
	GroupID: "misc",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(jailCmd).Standalone()

	rootCmd.AddCommand(jailCmd)

	carapace.Gen(jailCmd).PositionalCompletion(
		carapace.ActionValues(),
		zfs.ActionFilesystems(),
	)
}
