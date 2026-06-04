package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/shell"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "complete",
	Short: "Specify how arguments are to be completed by Readline",
	Long:  "https://www.gnu.org/software/bash/manual/html_node/Bash-Builtins.html#index-complete",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringS("A", "A", "", "action type for completion")
	rootCmd.Flags().StringS("C", "C", "", "command to execute for completion")
	rootCmd.Flags().BoolS("D", "D", false, "apply completions as default for commands without specific completion")
	rootCmd.Flags().BoolS("E", "E", false, "apply completions to empty commands")
	rootCmd.Flags().StringS("F", "F", "", "function to execute for completion")
	rootCmd.Flags().StringS("G", "G", "", "glob pattern for completion")
	rootCmd.Flags().BoolS("I", "I", false, "apply completions to the initial word")
	rootCmd.Flags().StringS("P", "P", "", "prefix to add to each completion")
	rootCmd.Flags().StringS("S", "S", "", "suffix to add to each completion")
	rootCmd.Flags().StringS("W", "W", "", "wordlist for completion")
	rootCmd.Flags().StringS("X", "X", "", "filter pattern to exclude from completions")
	rootCmd.Flags().BoolS("a", "a", false, "alias action")
	rootCmd.Flags().BoolS("b", "b", false, "builtin action")
	rootCmd.Flags().BoolS("c", "c", false, "command action")
	rootCmd.Flags().BoolS("d", "d", false, "directory action")
	rootCmd.Flags().BoolS("e", "e", false, "exported variable action")
	rootCmd.Flags().BoolS("f", "f", false, "file action")
	rootCmd.Flags().BoolS("g", "g", false, "group action")
	rootCmd.Flags().BoolS("j", "j", false, "job action")
	rootCmd.Flags().BoolS("k", "k", false, "keyword action")
	rootCmd.Flags().StringS("o", "o", "", "completion option")
	rootCmd.Flags().BoolS("p", "p", false, "print existing completion specifications in a reusable format")
	rootCmd.Flags().BoolS("r", "r", false, "remove a completion specification")
	rootCmd.Flags().BoolS("s", "s", false, "service action")
	rootCmd.Flags().BoolS("u", "u", false, "user action")
	rootCmd.Flags().BoolS("v", "v", false, "variable action")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"A": carapace.ActionValuesDescribed(
			"alias", "aliases",
			"arrayvar", "array variables",
			"binding", "readline key bindings",
			"builtin", "shell builtins",
			"command", "commands",
			"directory", "directories",
			"disabled", "disabled builtins",
			"enabled", "enabled builtins",
			"export", "exported variables",
			"file", "files",
			"function", "shell functions",
			"group", "groups",
			"helptopic", "help topics",
			"hostname", "hostnames",
			"job", "jobs",
			"keyword", "shell reserved words",
			"running", "running jobs",
			"service", "service names",
			"setopt", "shell options",
			"shopt", "shell options",
			"signal", "signals",
			"stopped", "stopped jobs",
			"user", "users",
			"variable", "shell variables",
		),
		"C": carapace.ActionFiles(),
		"F": shell.ActionFunctions(true),
		"G": carapace.ActionFiles(),
		"P": carapace.ActionFiles(),
		"S": carapace.ActionFiles(),
		"W": carapace.ActionValues(),
		"X": carapace.ActionValues(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		shell.ActionExecutables(),
	)
}
