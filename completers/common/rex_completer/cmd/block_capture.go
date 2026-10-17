package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/rex_completer/cmd/common"
	"github.com/spf13/cobra"
)

var blockCaptureCmd = &cobra.Command{
	Use:   "capture",
	Short: "Print a terminal's screen",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(blockCaptureCmd).Standalone()

	blockCaptureCmd.Flags().String("format", "text", "output FORMAT: text, html, or vt")
	blockCaptureCmd.Flags().Bool("trim", false, "remove trailing blank lines")
	blockCaptureCmd.Flags().Bool("unwrap", false, "join soft-wrapped lines")

	common.AddScopeFlags(blockCaptureCmd)

	blockCmd.AddCommand(blockCaptureCmd)

	carapace.Gen(blockCaptureCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("html", "text", "vt"),
	})
}
