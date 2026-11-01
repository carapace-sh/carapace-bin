package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dyld_usage",
	Short: "report dynamic linker activity in real-time",
	Long:  "https://man.freebsd.org/cgi/man.cgi?dyld_usage",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringS("R", "R", "", "Specify a raw trace file to process")
	rootCmd.Flags().BoolS("e", "e", false, "Exclude specified list of pids and commands")
	rootCmd.Flags().BoolS("h", "h", false, "Display usage information")
	rootCmd.Flags().BoolS("j", "j", false, "Display output in JSON format")
	rootCmd.Flags().StringS("t", "t", "", "Timeout in seconds")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"R": carapace.ActionFiles(),
	})
}
