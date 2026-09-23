package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "termux-restore",
	Short: "Restore the Termux prefix from a tar archive",
	Long:  "https://github.com/termux/termux-tools/blob/master/scripts/termux-restore.in",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("help", "h", false, "show usage and exit")
	rootCmd.Flags().BoolP("usage", "?", false, "show usage and exit")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
