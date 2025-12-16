package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/just"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var rootCmd = &cobra.Command{
	Use:   "just",
	Short: "Just a command runner",
	Long:  "https://github.com/casey/just",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute(opts ...func(cmd *cobra.Command)) error {
	for _, opt := range opts {
		opt(rootCmd)
	}
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("changelog", false, "Print changelog")
	rootCmd.Flags().Bool("check", false, "Run `--fmt` in 'check' mode. Exits with 0 if justfile is formatted correctly. Exits with 1")
	rootCmd.Flags().Bool("choose", false, "Select one or more recipes to run using a binary chooser. If `--chooser` is not passed the")
	rootCmd.Flags().String("chooser", "", "Override binary invoked by `--choose` [env: JUST_CHOOSER=]")
	rootCmd.Flags().Bool("clear-shell-args", false, "Clear shell arguments")
	rootCmd.Flags().String("color", "", "Print colorful output [env: JUST_COLOR=] [default: auto] [possible values: auto, always,")
	rootCmd.Flags().StringP("command", "c", "", "Run an arbitrary command with the working directory, `.env`, overrides, and exports set")
	rootCmd.Flags().String("command-color", "", "Echo recipe lines in <COMMAND-COLOR> [env: JUST_COMMAND_COLOR=] [possible values: black,")
	rootCmd.Flags().String("completions", "", "Print shell completion script for <SHELL> [possible values: bash, elvish, fish, nushell,")
	rootCmd.Flags().String("dotenv-filename", "", "Search for environment file named <DOTENV-FILENAME> instead of `.env`")
	rootCmd.Flags().StringP("dotenv-path", "E", "", "Load <DOTENV-PATH> as environment file instead of searching for one")
	rootCmd.Flags().BoolP("dry-run", "n", false, "Print what just would do without doing it [env: JUST_DRY_RUN=]")
	rootCmd.Flags().Bool("dump", false, "Print justfile")
	rootCmd.Flags().String("dump-format", "", "Dump justfile as <FORMAT> [env: JUST_DUMP_FORMAT=] [default: just] [possible values: json,")
	rootCmd.Flags().BoolP("edit", "e", false, "Edit justfile with editor given by $VISUAL or $EDITOR, falling back to `vim`")
	rootCmd.Flags().Bool("evaluate", false, "Evaluate and print all variables. If a variable name is given as an argument, only print")
	rootCmd.Flags().Bool("fmt", false, "Format and overwrite justfile")
	rootCmd.Flags().BoolP("global-justfile", "g", false, "Use global justfile")
	rootCmd.Flags().Bool("groups", false, "List recipe groups")
	rootCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.Flags().Bool("highlight", false, "Highlight echoed recipe lines in bold [env: JUST_HIGHLIGHT=]")
	rootCmd.Flags().Bool("init", false, "Initialize new justfile in project root")
	rootCmd.Flags().StringP("justfile", "f", "", "Use <JUSTFILE> as justfile [env: JUST_JUSTFILE=]")
	rootCmd.Flags().StringP("list", "l", "", "List available recipes")
	rootCmd.Flags().String("list-heading", "", "Print <TEXT> before list [env: JUST_LIST_HEADING=] [default: \"Available recipes:\n\"]")
	rootCmd.Flags().String("list-prefix", "", "Print <TEXT> before each list item [env: JUST_LIST_PREFIX=] [default: \"    \"]")
	rootCmd.Flags().Bool("list-submodules", false, "List recipes in submodules [env: JUST_LIST_SUBMODULES=]")
	rootCmd.Flags().Bool("man", false, "Print man page")
	rootCmd.Flags().Bool("no-aliases", false, "Don't show aliases in list [env: JUST_NO_ALIASES=]")
	rootCmd.Flags().Bool("no-deps", false, "Don't run recipe dependencies [env: JUST_NO_DEPS=]")
	rootCmd.Flags().Bool("no-dotenv", false, "Don't load `.env` file [env: JUST_NO_DOTENV=]")
	rootCmd.Flags().Bool("no-highlight", false, "Don't highlight echoed recipe lines in bold [env: JUST_NO_HIGHLIGHT=]")
	rootCmd.Flags().BoolP("quiet", "q", false, "Suppress all output [env: JUST_QUIET=]")
	rootCmd.Flags().StringSlice("set", nil, "Override <VARIABLE> with <VALUE>")
	rootCmd.Flags().String("shell", "", "Invoke <SHELL> to run recipes")
	rootCmd.Flags().String("shell-arg", "", "Invoke shell with <SHELL-ARG> as an argument")
	rootCmd.Flags().Bool("shell-command", false, "Invoke <COMMAND> with the shell used to run recipe lines and backticks")
	rootCmd.Flags().StringP("show", "s", "", "Show recipe at <PATH>")
	rootCmd.Flags().Bool("summary", false, "List names of available recipes")
	rootCmd.Flags().Bool("timestamp", false, "Print recipe command timestamps [env: JUST_TIMESTAMP=]")
	rootCmd.Flags().String("timestamp-format", "", "Timestamp format string [env: JUST_TIMESTAMP_FORMAT=] [default: %H:%M:%S]")
	rootCmd.Flags().BoolP("unsorted", "u", false, "Return list and summary entries in source order [env: JUST_UNSORTED=]")
	rootCmd.Flags().Bool("unstable", false, "Enable unstable features [env: JUST_UNSTABLE=]")
	rootCmd.Flags().Bool("variables", false, "List names of variables")
	rootCmd.Flags().CountP("verbose", "v", "Use verbose output [env: JUST_VERBOSE=]")
	rootCmd.Flags().BoolP("version", "V", false, "Print version")
	rootCmd.Flags().StringP("working-directory", "d", "", "Use <WORKING-DIRECTORY> as working directory. --justfile must also be set [env:")
	rootCmd.Flags().Bool("yes", false, "Automatically confirm all recipes. [env: JUST_YES=]")

	rootCmd.Flag("set").Nargs = 2

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"chooser": carapace.Batch(
			carapace.ActionExecutables(),
			carapace.ActionFiles(),
		).ToA(),
		"color":         carapace.ActionValues("auto", "always", "never").StyleF(style.ForKeyword),
		"command-color": carapace.ActionValues("auto", "always", "never").StyleF(style.ForKeyword),
		"completions":   carapace.ActionValues("bash", "elvish", "fish", "nushell", "powershell", "zsh"),
		"dotenv-path":   carapace.ActionDirectories(),
		"dump-format":   carapace.ActionValues("just", "json"),
		"justfile":      carapace.ActionFiles(),
		"list":          carapace.ActionFiles(),
		"set": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return just.ActionVariables(rootCmd.Flag("justfile").Value.String())
			default:
				return carapace.ActionValues()
			}
		}),
		"shell": os.ActionShells(),
		"show": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return just.ActionRecipes(rootCmd.Flag("justfile").Value.String())
		}),
		"working-directory": carapace.ActionDirectories(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return just.ActionRecipes(rootCmd.Flag("justfile").Value.String()).FilterArgs()
		}),
	)

	carapace.Gen(rootCmd).PreInvoke(func(cmd *cobra.Command, flag *pflag.Flag, action carapace.Action) carapace.Action {
		if f := rootCmd.Flag("working-directory"); f.Changed {
			return action.Chdir(f.Value.String())
		}
		return action
	})
}
