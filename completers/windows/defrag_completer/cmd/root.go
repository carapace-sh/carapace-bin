package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "defrag",
	Short: "locate and consolidate fragmented files on local volumes",
	Long:  "https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/defrag",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("a", "a", false, "analyze only, do not defragment")
	rootCmd.Flags().BoolP("h", "h", false, "run at normal priority")
	rootCmd.Flags().BoolP("m", "m", false, "run on each volume in parallel")
	rootCmd.Flags().BoolP("v", "v", false, "print verbose output")
}
