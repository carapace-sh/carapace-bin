package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/rbenv_completer/cmd/action"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "rbenv",
	Short: "Manage your application's Ruby environment",
	Long:  "https://github.com/rbenv/rbenv",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

var commandsCmd = &cobra.Command{
	Use:   "commands",
	Short: "List all available rbenv commands",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var completionsCmd = &cobra.Command{
	Use:   "completions",
	Short: "Provide completion candidates for a command",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "Run an executable with the selected Ruby version",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var globalCmd = &cobra.Command{
	Use:   "global",
	Short: "Set or show the global Ruby version",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Display help for a command",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var hooksCmd = &cobra.Command{
	Use:   "hooks",
	Short: "List hook scripts for a given rbenv command",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Configure the shell environment for rbenv",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var localCmd = &cobra.Command{
	Use:   "local",
	Short: "Set or show the local application-specific Ruby version",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var prefixCmd = &cobra.Command{
	Use:   "prefix",
	Short: "Display prefix for a Ruby version",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var rehashCmd = &cobra.Command{
	Use:   "rehash",
	Short: "Regenerate rbenv shims",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var rootSubCmd = &cobra.Command{
	Use:   "root",
	Short: "Display the root directory where versions and shims are kept",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var shellCmd = &cobra.Command{
	Use:   "shell",
	Short: "Set or show the shell-specific Ruby version",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var shimsCmd = &cobra.Command{
	Use:   "shims",
	Short: "List existing rbenv shims",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the current Ruby version and its origin",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var versionFileCmd = &cobra.Command{
	Use:   "version-file",
	Short: "Detect the file that sets the current rbenv version",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var versionFileReadCmd = &cobra.Command{
	Use:   "version-file-read",
	Short: "Read a version file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var versionFileWriteCmd = &cobra.Command{
	Use:   "version-file-write",
	Short: "Write a version file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var versionNameCmd = &cobra.Command{
	Use:   "version-name",
	Short: "Show the current Ruby version",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var versionOriginCmd = &cobra.Command{
	Use:   "version-origin",
	Short: "Explain how the current Ruby version is set",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var versionsCmd = &cobra.Command{
	Use:   "versions",
	Short: "List installed Ruby versions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var whenceCmd = &cobra.Command{
	Use:   "whence",
	Short: "List all Ruby versions that contain the given executable",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var whichCmd = &cobra.Command{
	Use:   "which",
	Short: "Display the full path to an executable",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("version", false, "Display the version of rbenv")

	carapace.Gen(commandsCmd).Standalone()
	carapace.Gen(completionsCmd).Standalone()
	carapace.Gen(execCmd).Standalone()
	carapace.Gen(globalCmd).Standalone()
	carapace.Gen(helpCmd).Standalone()
	carapace.Gen(hooksCmd).Standalone()
	carapace.Gen(initCmd).Standalone()
	carapace.Gen(localCmd).Standalone()
	carapace.Gen(prefixCmd).Standalone()
	carapace.Gen(rehashCmd).Standalone()
	carapace.Gen(rootSubCmd).Standalone()
	carapace.Gen(shellCmd).Standalone()
	carapace.Gen(shimsCmd).Standalone()
	carapace.Gen(versionCmd).Standalone()
	carapace.Gen(versionFileCmd).Standalone()
	carapace.Gen(versionFileReadCmd).Standalone()
	carapace.Gen(versionFileWriteCmd).Standalone()
	carapace.Gen(versionNameCmd).Standalone()
	carapace.Gen(versionOriginCmd).Standalone()
	carapace.Gen(versionsCmd).Standalone()
	carapace.Gen(whenceCmd).Standalone()
	carapace.Gen(whichCmd).Standalone()

	rootCmd.AddCommand(commandsCmd)
	rootCmd.AddCommand(completionsCmd)
	rootCmd.AddCommand(execCmd)
	rootCmd.AddCommand(globalCmd)
	rootCmd.AddCommand(helpCmd)
	rootCmd.AddCommand(hooksCmd)
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(localCmd)
	rootCmd.AddCommand(prefixCmd)
	rootCmd.AddCommand(rehashCmd)
	rootCmd.AddCommand(rootSubCmd)
	rootCmd.AddCommand(shellCmd)
	rootCmd.AddCommand(shimsCmd)
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(versionFileCmd)
	rootCmd.AddCommand(versionFileReadCmd)
	rootCmd.AddCommand(versionFileWriteCmd)
	rootCmd.AddCommand(versionNameCmd)
	rootCmd.AddCommand(versionOriginCmd)
	rootCmd.AddCommand(versionsCmd)
	rootCmd.AddCommand(whenceCmd)
	rootCmd.AddCommand(whichCmd)

	commandsCmd.Flags().Bool("no-sh", false, "")
	commandsCmd.Flags().Bool("sh", false, "")

	helpCmd.Flags().Bool("usage", false, "")

	shimsCmd.Flags().Bool("short", false, "")

	versionsCmd.Flags().Bool("bare", false, "")
	versionsCmd.Flags().Bool("skip-aliases", false, "")

	whenceCmd.Flags().Bool("path", false, "")

	carapace.Gen(completionsCmd).PositionalCompletion(
		action.ActionCommands(),
	)

	carapace.Gen(execCmd).PositionalCompletion(
		action.ActionShims(),
	)

	carapace.Gen(globalCmd).PositionalCompletion(
		action.ActionVersionsAndSystem(),
	)

	carapace.Gen(helpCmd).PositionalCompletion(
		action.ActionCommands(),
	)

	carapace.Gen(hooksCmd).PositionalCompletion(
		action.ActionHookCommands(),
	)

	carapace.Gen(initCmd).PositionalCompletion(
		action.ActionShells(),
	)

	carapace.Gen(localCmd).PositionalCompletion(
		action.ActionVersionsAndSystem(),
	)

	carapace.Gen(prefixCmd).PositionalCompletion(
		action.ActionVersionsAndSystem(),
	)

	carapace.Gen(shellCmd).PositionalCompletion(
		action.ActionVersionsAndSystem(),
	)

	carapace.Gen(versionFileCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)

	carapace.Gen(versionFileReadCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)

	carapace.Gen(versionFileWriteCmd).PositionalCompletion(
		carapace.ActionFiles(),
		action.ActionVersionsAndSystem(),
	)

	carapace.Gen(whenceCmd).PositionalCompletion(
		action.ActionShims(),
	)

	carapace.Gen(whichCmd).PositionalCompletion(
		action.ActionShims(),
	)

}
