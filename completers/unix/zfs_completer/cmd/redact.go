package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zfs"
	"github.com/spf13/cobra"
)

var redactCmd = &cobra.Command{
	Use:     "redact snapshot redaction_bookmark redaction_snapshot...",
	Short:   "create a redaction bookmark",
	GroupID: "send",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redactCmd).Standalone()

	rootCmd.AddCommand(redactCmd)

	carapace.Gen(redactCmd).PositionalCompletion(
		zfs.ActionSnapshots(),
		carapace.ActionValues(),
	)

	carapace.Gen(redactCmd).PositionalAnyCompletion(
		zfs.ActionSnapshots(),
	)
}
