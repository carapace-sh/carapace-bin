package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "termux-info",
	Short: "Print system and debug information",
	Long:  "https://github.com/termux/termux-tools/blob/master/scripts/termux-info.in",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("help", "h", false, "show usage and exit")
	rootCmd.Flags().Bool("no-set-clipboard", false, "do not copy output to the clipboard")
}
