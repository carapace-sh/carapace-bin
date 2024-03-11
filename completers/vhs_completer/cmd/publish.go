package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var publishCmd = &cobra.Command{
	Use:   "publish <gif>",
	Short: "Publish your GIF to vhs.charm.sh and get a shareable URL",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(publishCmd).Standalone()

	rootCmd.AddCommand(publishCmd)

	carapace.Gen(publishCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
