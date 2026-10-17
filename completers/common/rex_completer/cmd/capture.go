package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/rex_completer/cmd/common"
	"github.com/spf13/cobra"
)

var captureCmd = &cobra.Command{
	Use:   "capture",
	Short: "Print a terminal's screen",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(captureCmd).Standalone()

	captureCmd.Flags().String("format", "text", "output FORMAT: text, html, or vt")
	captureCmd.Flags().Bool("trim", false, "remove trailing blank lines")
	captureCmd.Flags().Bool("unwrap", false, "join soft-wrapped lines")

	common.AddScopeFlags(captureCmd)

	rootCmd.AddCommand(captureCmd)

	carapace.Gen(captureCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("html", "text", "vt"),
	})
}
