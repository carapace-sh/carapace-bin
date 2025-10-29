package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/d2"
	"github.com/spf13/cobra"
)

var layoutCmd = &cobra.Command{
	Use:   "layout [name]",
	Short: "Lists available layout engine options with short help",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(layoutCmd).Standalone()

	rootCmd.AddCommand(layoutCmd)

	carapace.Gen(layoutCmd).PositionalCompletion(
		d2.ActionLayouts(),
	)
}
