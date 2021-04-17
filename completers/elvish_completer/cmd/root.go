package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "elvish",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().String("bin", "", "path to the elvish binary")
	rootCmd.Flags().Bool("buildinfo", false, "show build info and quit")
	rootCmd.Flags().BoolS("c", "c", false, "take first argument as code to execute")
	rootCmd.Flags().Bool("compileonly", false, "Parse/Compile but do not execute")
	rootCmd.Flags().String("cpuprofile", "", "write cpu profile to file")
	rootCmd.Flags().Bool("daemon", false, "run daemon instead of shell")
	rootCmd.Flags().String("db", "", "path to the database")
	rootCmd.Flags().Bool("help", false, "show usage help and quit")
	rootCmd.Flags().Bool("json", false, "show output in JSON. Useful with --buildinfo.")
	rootCmd.Flags().String("log", "", "a file to write debug log to")
	rootCmd.Flags().String("logprefix", "", "the prefix for the daemon log file")
	rootCmd.Flags().Bool("norc", false, "run elvish without invoking rc.elv")
	rootCmd.Flags().String("port", "", "the port of the web backend (default 3171)")
	rootCmd.Flags().Bool("show-deprecations", false, "whether to show deprecations")
	rootCmd.Flags().String("sock", "", "path to the daemon socket")
	rootCmd.Flags().Bool("version", false, "show version and quit")
	rootCmd.Flags().Bool("web", false, "run backend of web interface")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"bin":        carapace.ActionFiles(),
		"cpuprofile": carapace.ActionFiles(),
		"db":         carapace.ActionFiles(),
		"log":        carapace.ActionFiles(),
		"sock":       carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if rootCmd.Flag("c").Changed {
				return carapace.ActionValues()
			} else {
				return carapace.ActionFiles()
			}
		}),
	)

}
