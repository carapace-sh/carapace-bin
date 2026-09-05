package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show commit details and associated metadata",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(showCmd).Standalone()

	showCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(showCmd)

	carapace.Gen(showCmd).PositionalCompletion(
		git.ActionRefs(git.RefOption{}.Default()),
	)
}
