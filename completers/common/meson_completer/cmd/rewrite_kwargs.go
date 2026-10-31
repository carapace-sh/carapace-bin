package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rewrite_kwargsCmd = &cobra.Command{
	Use:   "kwargs",
	Short: "Modify keyword arguments",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rewrite_kwargsCmd).Standalone()

	applyCommonFlagCompletion(rewrite_kwargsCmd)
	rewriteCmd.AddCommand(rewrite_kwargsCmd)
}
