package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "termux-wake-unlock",
	Short: "Release a partial wake lock",
	Long:  "https://github.com/termux/termux-tools/blob/master/scripts/termux-wake-unlock.in",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

}
