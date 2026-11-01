package cmd

import (
	"github.com/carapace-sh/carapace"
	android "github.com/carapace-sh/carapace-bin/pkg/actions/tools/android"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "termux-open-url",
	Short: "Open a URL for viewing",
	Long:  "https://github.com/termux/termux-tools/blob/master/scripts/termux-open-url.in",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValues(), // URL
		android.ActionPackages(),
	)
}
