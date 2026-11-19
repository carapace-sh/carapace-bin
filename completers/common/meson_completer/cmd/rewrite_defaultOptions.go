package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rewrite_defaultOptionsCmd = &cobra.Command{
	Use:   "default-options",
	Short: "Modify the project default options",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rewrite_defaultOptionsCmd).Standalone()

	applyCommonFlagCompletion(rewrite_defaultOptionsCmd)
	rewriteCmd.AddCommand(rewrite_defaultOptionsCmd)
}
