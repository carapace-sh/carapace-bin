package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/scoop"
	"github.com/spf13/cobra"
)

var catCmd = &cobra.Command{
	Use:   "cat",
	Short: "show content of specified manifest",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(catCmd).Standalone()
	rootCmd.AddCommand(catCmd)

	carapace.Gen(catCmd).PositionalCompletion(
		scoop.ActionAvailableApps(),
	)
}
