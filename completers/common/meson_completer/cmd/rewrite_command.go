package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rewrite_commandCmd = &cobra.Command{
	Use:   "command",
	Short: "Execute a JSON array of commands",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rewrite_commandCmd).Standalone()

	applyCommonFlagCompletion(rewrite_commandCmd)
	rewriteCmd.AddCommand(rewrite_commandCmd)
}
