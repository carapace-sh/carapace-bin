package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "elvish",
	Short: "expressive programming language and a versatile interactive shell",
	Long:  "https://elv.sh/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	carapace.Override(carapace.Opts{
		LongShorthand: true,
	})
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("buildinfo", false, "Output information about the Elvish build and quit")
	rootCmd.Flags().Bool("c", false, "Treat the first argument as code to execute")
	rootCmd.Flags().Bool("compileonly", false, "Parse and compile Elvish code without executing it")
	rootCmd.Flags().Bool("daemon", false, "[internal flag] Run the storage daemon instead of an Elvish shell")
	rootCmd.Flags().String("db", "", "[internal flag] Path to the database file")
	rootCmd.Flags().String("deprecation-level", "", "Show warnings for all features deprecated as of version 0.X (default 18)")
	rootCmd.Flags().Bool("help", false, "Show usage help and quit")
	rootCmd.Flags().Bool("i", false, "A no-op flag, introduced for POSIX compatibility")
	rootCmd.Flags().Bool("json", false, "Show the output from -buildinfo, -compileonly or -version in JSON")
	rootCmd.Flags().String("log", "", "Path to a file to write debug logs")
	rootCmd.Flags().Bool("lsp", false, "Run the builtin language server")
	rootCmd.Flags().Bool("norc", false, "Don't read the RC file when running interactively")
	rootCmd.Flags().String("rc", "", "Path to the RC file when running interactively")
	rootCmd.Flags().String("sock", "", "[internal flag] Path to the daemon's UNIX socket")
	rootCmd.Flags().Bool("version", false, "Output the Elvish version and quit")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"db":   carapace.ActionFiles(),
		"log":  carapace.ActionFiles(),
		"rc":   carapace.ActionFiles(),
		"sock": carapace.ActionFiles(),
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
