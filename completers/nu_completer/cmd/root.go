package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "nu",
	Short: "Nushell",
	Long:  "https://www.nushell.sh/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("commands", "c", "", "run the given commands and then exit")
	rootCmd.Flags().String("config", "", "start with an alternate config file")
	rootCmd.Flags().String("env-config", "", "start with an alternate environment config file")
	rootCmd.Flags().StringP("execute", "e", "", "run the given commands and then enter an interactive shell")
	rootCmd.Flags().BoolP("help", "h", false, "Display the help message for this command")
	rootCmd.Flags().Bool("ide-ast", false, "generate the ast on the given source")
	rootCmd.Flags().String("ide-check", "", "run a diagnostic check on the given source and limit number of errors returned to provided number")
	rootCmd.Flags().String("ide-complete", "", "list completions for the item at the given position")
	rootCmd.Flags().String("ide-goto-def", "", "go to the definition of the item at the given position")
	rootCmd.Flags().String("ide-hover", "", "give information about the item at the given position")
	rootCmd.Flags().StringP("include-path", "I", "", "set the NU_LIB_DIRS for the given script (semicolon-delimited)")
	rootCmd.Flags().BoolP("interactive", "i", false, "start as an interactive shell")
	rootCmd.Flags().String("log-level", "", "log level for diagnostic logs (error, warn, info, debug, trace). Off by default")
	rootCmd.Flags().String("log-target", "", "set the target for the log to output. stdout, stderr(default), mixed or file")
	rootCmd.Flags().BoolP("login", "l", false, "start as a login shell")
	rootCmd.Flags().BoolP("no-config-file", "n", false, "start with no config file and no env file")
	rootCmd.Flags().Bool("no-std-lib", false, "start with no standard library")
	rootCmd.Flags().String("plugin-config", "", "start with an alternate plugin signature file")
	rootCmd.Flags().Bool("stdin", false, "redirect standard input to a command (with `-c`) or a script file")
	rootCmd.Flags().StringP("table-mode", "m", "", "the table mode to use. rounded is default.")
	rootCmd.Flags().String("testbin", "", "run internal test binary")
	rootCmd.Flags().StringP("threads", "t", "", "threads to use for parallel commands")
	rootCmd.Flags().BoolP("version", "v", false, "print the version")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"commands":      bridge.ActionCarapaceBin().SplitP(),
		"config":        carapace.ActionFiles(),
		"env-config":    carapace.ActionFiles(),
		"execute":       bridge.ActionCarapaceBin().SplitP(),
		"include-path":  carapace.ActionDirectories().List(";"),
		"log-level":     carapace.ActionValues("error", "warn", "info", "debug", "trace").StyleF(style.ForLogLevel),
		"log-target":    carapace.ActionValues("stdout", "stderr", "mixed", "file"),
		"plugin-config": carapace.ActionFiles(),
		"table-mode":    carapace.ActionValues("rounded", "basic", "compact", "compact_double", "light", "thin", "with_love", "reinforced", "heavy", "none", "other"),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
