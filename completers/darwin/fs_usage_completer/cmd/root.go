package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "fs_usage",
	Short: "file system usage trace",
	Long:  "https://keith.github.io/xcode-manpages/fs_usage.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("bootcache", "b", false, "Annotate disk I/O events with BootCache info")
	rootCmd.Flags().StringP("endtime", "E", "", "End time for trace processing")
	rootCmd.Flags().BoolP("exclude", "e", false, "Exclude sampling data from output")
	rootCmd.Flags().BoolP("filter", "f", false, "Turn on output filtering based on the filter file")
	rootCmd.Flags().StringP("rawfile", "R", "", "Process raw trace data from the specified file")
	rootCmd.Flags().StringP("starttime", "S", "", "Start time for trace processing")
	rootCmd.Flags().StringP("time", "t", "", "Set timeout in seconds")
	rootCmd.Flags().BoolP("trim", "F", false, "Trim pathnames from the right side")
	rootCmd.Flags().BoolP("wide", "w", false, "Force wider, more detailed output")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"rawfile": carapace.ActionFiles(),
	})
}
