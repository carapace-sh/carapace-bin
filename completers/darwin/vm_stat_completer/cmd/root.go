package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "vm_stat",
	Short: "show Mach virtual memory statistics",
	Long:  "https://keith.github.io/xcode-manpages/vm_stat.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("count", "c", "", "Terminate after count intervals")
}
