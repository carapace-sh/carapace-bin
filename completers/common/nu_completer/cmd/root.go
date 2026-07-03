package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
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
	rootCmd.Flags().String("error-style", "", "the error style to use (fancy or plain). default: fancy")
	rootCmd.Flags().StringP("execute", "e", "", "run the given commands and then enter an interactive shell")
	rootCmd.Flags().String("experimental-options", "", "enable or disable experimental options, use \"all\" to set all active options")
	rootCmd.Flags().BoolP("help", "h", false, "Display the help message for this command")
	rootCmd.Flags().Bool("ide-ast", false, "generate the ast on the given source")
	rootCmd.Flags().String("ide-check", "", "run a diagnostic check on the given source and limit number of errors returned to provided number")
	rootCmd.Flags().String("ide-complete", "", "list completions for the item at the given position")
	rootCmd.Flags().String("ide-goto-def", "", "go to the definition of the item at the given position")
	rootCmd.Flags().String("ide-hover", "", "give information about the item at the given position")
	rootCmd.Flags().StringP("include-path", "I", "", "set the NU_LIB_DIRS for the given script (semicolon-delimited)")
	rootCmd.Flags().BoolP("interactive", "i", false, "start as an interactive shell")
	rootCmd.Flags().String("log-exclude", "", "set the Rust module prefixes to exclude from the log output")
	rootCmd.Flags().String("log-file", "", "specify a custom log file path (requires --log-target file and --log-level <level>)")
	rootCmd.Flags().String("log-include", "", "set the Rust module prefixes to include from the log output")
	rootCmd.Flags().String("log-level", "", "log level for diagnostic logs (error, warn, info, debug, trace). Off by default")
	rootCmd.Flags().String("log-target", "", "set the target for the log to output. stdout, stderr(default), mixed or file")
	rootCmd.Flags().BoolP("login", "l", false, "start as a login shell")
	rootCmd.Flags().Bool("lsp", false, "start nu's language server protocol")
	rootCmd.Flags().Bool("mcp", false, "start nu's model context protocol server")
	rootCmd.Flags().String("mcp-port", "", "port for MCP HTTP transport (default 8080)")
	rootCmd.Flags().String("mcp-transport", "", "transport to use for MCP server (stdio or http)")
	rootCmd.Flags().BoolP("no-config-file", "n", false, "start with no config file and no env file")
	rootCmd.Flags().Bool("no-history", false, "disable reading and writing to command history")
	rootCmd.Flags().Bool("no-newline", false, "print the result for --commands(-c) without a newline")
	rootCmd.Flags().Bool("no-std-lib", false, "start with no standard library")
	rootCmd.Flags().String("plugin-config", "", "start with an alternate plugin registry file")
	rootCmd.Flags().StringSlice("plugins", nil, "list of plugin executable files to load (full paths), separately from the registry file")
	rootCmd.Flags().Bool("stdin", false, "redirect standard input to a command (with `-c`) or a script file")
	rootCmd.Flags().StringP("table-mode", "m", "", "the table mode to use. rounded is default.")
	rootCmd.Flags().String("testbin", "", "run internal test binary")
	rootCmd.Flags().StringP("threads", "t", "", "threads to use for parallel commands")
	rootCmd.Flags().BoolP("version", "v", false, "print the version")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"commands":      bridge.ActionCarapaceBin().SplitP(),
		"config":        carapace.ActionFiles(),
		"env-config":    carapace.ActionFiles(),
		"error-style":   carapace.ActionValues("fancy", "plain"),
		"execute":       bridge.ActionCarapaceBin().SplitP(),
		"include-path":  carapace.ActionDirectories().List(";"),
		"log-file":      carapace.ActionFiles(),
		"log-level":     carapace.ActionValues("error", "warn", "info", "debug", "trace").StyleF(style.ForLogLevel),
		"log-target":    carapace.ActionValues("stdout", "stderr", "mixed", "file"),
		"mcp-port":      net.ActionPorts(),
		"mcp-transport": carapace.ActionValues("stdio", "http"),
		"plugin-config": carapace.ActionFiles(),
		"plugins":       carapace.ActionFiles(),
		"table-mode":    carapace.ActionValues("rounded", "basic", "compact", "compact_double", "light", "thin", "with_love", "reinforced", "heavy", "none", "other"),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
