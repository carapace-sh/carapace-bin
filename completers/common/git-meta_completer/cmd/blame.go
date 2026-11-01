package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var blameCmd = &cobra.Command{
	Use:   "blame",
	Short: "Show file blame grouped by pull request metadata",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(blameCmd).Standalone()

	blameCmd.Flags().BoolP("help", "h", false, "Print help")
	blameCmd.Flags().Bool("json", false, "Output compact JSON with line ranges and PR metadata")
	blameCmd.Flags().Bool("porcelain", false, "Output machine-readable JSON")
	blameCmd.Flags().String("rev", "", "Revision to blame from (default: HEAD)")
	rootCmd.AddCommand(blameCmd)

	carapace.Gen(blameCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)

	carapace.Gen(blameCmd).FlagCompletion(carapace.ActionMap{
		"rev": git.ActionRefs(git.RefOption{}.Default()),
	})
}
