package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rewrite_cmdCmd = &cobra.Command{
	Use:   "cmd",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rewrite_cmdCmd).Standalone()

	applyCommonFlagCompletion(rewrite_cmdCmd)
	rewriteCmd.AddCommand(rewrite_cmdCmd)
}
