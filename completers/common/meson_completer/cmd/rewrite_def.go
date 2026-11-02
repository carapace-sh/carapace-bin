package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rewrite_defCmd = &cobra.Command{
	Use:   "def",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rewrite_defCmd).Standalone()

	applyCommonFlagCompletion(rewrite_defCmd)
	rewriteCmd.AddCommand(rewrite_defCmd)
}
