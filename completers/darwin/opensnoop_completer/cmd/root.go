package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "opensnoop",
	Short: "Trace file open operations",
	Long:  "https://keith.github.io/xcode-manpages/opensnoop.1m.html",
	Run:   func(*cobra.Command, []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("all", "a", false, "print most data")
	rootCmd.Flags().BoolP("alldata", "A", false, "dump all data, space delimited")
	rootCmd.Flags().BoolP("args", "g", false, "print command arguments")
	rootCmd.Flags().BoolP("cwd", "c", false, "print cwd of process")
	rootCmd.Flags().BoolP("errno", "e", false, "print errno value")
	rootCmd.Flags().BoolP("failed", "x", false, "only print failed opens")
	rootCmd.Flags().BoolP("flags", "F", false, "print the flags passed to open")
	rootCmd.Flags().StringP("name", "n", "", "process name to snoop")
	rootCmd.Flags().StringP("pathname", "f", "", "pathname name to snoop")
	rootCmd.Flags().StringP("pid", "p", "", "process ID to snoop")
	rootCmd.Flags().BoolP("stack", "t", false, "print user stack trace")
	rootCmd.Flags().BoolP("start", "s", false, "print start time, us")
	rootCmd.Flags().BoolP("verbose", "v", false, "print start time, string")
	rootCmd.Flags().BoolP("zone", "Z", false, "print zonename")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"pathname": carapace.ActionFiles(),
	})
}
