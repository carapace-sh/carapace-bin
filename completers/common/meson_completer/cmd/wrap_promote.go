package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wrap_promoteCmd = &cobra.Command{
	Use:   "promote",
	Short: "bring a subsubproject up to the master project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wrap_promoteCmd).Standalone()

	applyCommonFlagCompletion(wrap_promoteCmd)
	wrapCmd.AddCommand(wrap_promoteCmd)
}
