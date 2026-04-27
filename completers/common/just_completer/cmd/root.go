package cmd

import (
	"path/filepath"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/just"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
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

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()
	rootCmd.Flags().SetInterspersed(false)

	rootCmd.Flags().String("alias-style", "", "Set list command alias display style")
	rootCmd.Flags().Bool("allow-missing", false, "Ignore missing recipe and module errors")
	rootCmd.Flags().String("ceiling", "", "Do not ascend above <CEILING> directory when searching for a justfile.")
	rootCmd.Flags().Bool("changelog", false, "Print changelog")
	rootCmd.Flags().Bool("check", false, "Run `--fmt` in 'check' mode. Exits with 0 if justfile is formatted correctly. Exits with 1 and prints a diff if formatting is required.")
	rootCmd.Flags().Bool("choose", false, "Select one or more recipes to run using a binary chooser. If `--chooser` is not passed the chooser defaults to the value of $JUST_CHOOSER, falling back to `fzf`")
	rootCmd.Flags().String("chooser", "", "Override binary invoked by `--choose`")
	rootCmd.Flags().Bool("clear-shell-args", false, "Clear shell arguments")
	rootCmd.Flags().String("color", "", "Print colorful output")
	rootCmd.Flags().StringSliceP("command", "c", nil, "Run an arbitrary command with the working directory, `.env`, overrides, and exports set")
	rootCmd.Flags().String("command-color", "", "Echo recipe lines in <COMMAND-COLOR>")
	rootCmd.Flags().Bool("complete-aliases", false, "Auto-complete recipe aliases")
	rootCmd.Flags().String("completions", "", "Print shell completion script for <SHELL>")
	rootCmd.Flags().String("cygpath", "", "Use binary at <CYGPATH> to convert between unix and Windows paths.")
	rootCmd.Flags().String("dotenv-filename", "", "Search for environment file named <DOTENV-FILENAME> instead of `.env`")
	rootCmd.Flags().StringP("dotenv-path", "E", "", "Load <DOTENV-PATH> as environment file instead of searching for one")
	rootCmd.Flags().BoolP("dry-run", "n", false, "Print what just would do without doing it")
	rootCmd.Flags().Bool("dump", false, "Print justfile")
	rootCmd.Flags().String("dump-format", "", "Dump justfile as <FORMAT>")
	rootCmd.Flags().BoolP("edit", "e", false, "Edit justfile with editor given by $VISUAL or $EDITOR, falling back to `vim`")
	rootCmd.Flags().Bool("eval", false, "Evaluate and print all variables. If a variable name is given as an argument, only print that variable's value.")
	rootCmd.Flags().Bool("evaluate", false, "Evaluate and print all variables. If a variable name is given as an argument, only print that variable's value.")
	rootCmd.Flags().String("evaluate-format", "", "Print evaluated variables in <FORMAT>")
	rootCmd.Flags().Bool("explain", false, "Print recipe doc comment before running it")
	rootCmd.Flags().Bool("fmt", false, "Format and overwrite justfile")
	rootCmd.Flags().Bool("format", false, "Format and overwrite justfile")
	rootCmd.Flags().BoolP("global-justfile", "g", false, "Use global justfile")
	rootCmd.Flags().StringSlice("group", nil, "Only list recipes in <GROUP>")
	rootCmd.Flags().Bool("groups", false, "List recipe groups")
	rootCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.Flags().Bool("highlight", false, "Highlight echoed recipe lines in bold")
	rootCmd.Flags().String("indentation", "", "Indent recipes bodies with <INDENTATION>")
	rootCmd.Flags().Bool("init", false, "Initialize new justfile in project root")
	rootCmd.Flags().Bool("initialize", false, "Initialize new justfile in project root")
	rootCmd.Flags().Bool("json", false, "Print justfile as JSON")
	rootCmd.Flags().StringP("justfile", "f", "", "Use <JUSTFILE> as justfile")
	rootCmd.Flags().StringSlice("justfile-name", nil, "Search for justfile named <NAME>")
	rootCmd.Flags().StringSliceP("list", "l", nil, "List available recipes in <MODULE> or root if omitted")
	rootCmd.Flags().String("list-heading", "", "Print <TEXT> before list")
	rootCmd.Flags().String("list-prefix", "", "Print <TEXT> before each list item")
	rootCmd.Flags().Bool("list-submodules", false, "List recipes in submodules")
	rootCmd.Flags().Bool("man", false, "Print man page")
	rootCmd.Flags().Bool("no-aliases", false, "Don't show aliases in list")
	rootCmd.Flags().Bool("no-dependencies", false, "Don't run recipe dependencies")
	rootCmd.Flags().Bool("no-deps", false, "Don't run recipe dependencies")
	rootCmd.Flags().Bool("no-dotenv", false, "Don't load `.env` file")
	rootCmd.Flags().Bool("no-highlight", false, "Don't highlight echoed recipe lines in bold")
	rootCmd.Flags().Bool("one", false, "Forbid multiple recipes from being invoked on the command line")
	rootCmd.Flags().BoolP("quiet", "q", false, "Suppress all output")
	rootCmd.Flags().StringSlice("set", nil, "Override <VARIABLE> with <VALUE>")
	rootCmd.Flags().String("shell", "", "Invoke <SHELL> to run recipes")
	rootCmd.Flags().StringSlice("shell-arg", nil, "Invoke shell with <SHELL-ARG> as an argument")
	rootCmd.Flags().Bool("shell-command", false, "Invoke <COMMAND> with the shell used to run recipe lines and backticks")
	rootCmd.Flags().StringSliceP("show", "s", nil, "Show recipe at <PATH>")
	rootCmd.Flags().Bool("summary", false, "List names of available recipes")
	rootCmd.Flags().String("tempdir", "", "Save temporary files to <TEMPDIR>.")
	rootCmd.Flags().Bool("time", false, "Print recipe execution time")
	rootCmd.Flags().Bool("timestamp", false, "Print recipe command timestamps")
	rootCmd.Flags().String("timestamp-format", "", "Timestamp format string")
	rootCmd.Flags().BoolP("unsorted", "u", false, "Return list and summary entries in source order")
	rootCmd.Flags().Bool("unstable", false, "Enable unstable features")
	rootCmd.Flags().StringSlice("usage", nil, "Print recipe usage information")
	rootCmd.Flags().Bool("variables", false, "List names of variables")
	rootCmd.Flags().CountP("verbose", "v", "Use verbose output")
	rootCmd.Flags().BoolP("version", "V", false, "Print version")
	rootCmd.Flags().StringP("working-directory", "d", "", "Use <WORKING-DIRECTORY> as working directory. --justfile must also be set")
	rootCmd.Flags().Bool("yes", false, "Automatically confirm all recipes.")
	rootCmd.Flag("eval").Hidden = true
	rootCmd.Flag("format").Hidden = true
	rootCmd.Flag("initialize").Hidden = true
	rootCmd.Flag("no-dependencies").Hidden = true

	rootCmd.Flag("set").Nargs = 2

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"alias-style": carapace.ActionValues("left", "right", "separate"),
		"ceiling":     carapace.ActionValues(), // TODO whats ceiling
		"chooser": carapace.Batch(
			carapace.ActionExecutables(),
			carapace.ActionFiles(),
		).ToA(),
		"color":           carapace.ActionValues("auto", "always", "never").StyleF(style.ForKeyword),
		"command":         bridge.ActionCarapaceBin().Split(),
		"command-color":   carapace.ActionValues("auto", "always", "never").StyleF(style.ForKeyword),
		"completions":     carapace.ActionValues("bash", "elvish", "fish", "nushell", "powershell", "zsh"),
		"cygpath":         carapace.ActionDirectories(),
		"dotenv-path":     carapace.ActionDirectories(),
		"dump-format":     carapace.ActionValues("just", "json"),
		"evaluate-format": carapace.ActionValues("just", "shell"),
		"group": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return just.ActionGroups(rootCmd.Flag("justfile").Value.String())
		}),
		"justfile": carapace.ActionFiles(),
		"list":     carapace.ActionFiles(),
		"set": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return just.ActionVariables(rootCmd.Flag("justfile").Value.String())
			default:
				return carapace.ActionValues()
			}
		}),
		"shell": os.ActionShells(),
		"shell-arg": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if f := rootCmd.Flag("shell"); f.Changed {
				shell := filepath.Base(f.Value.String())
				return bridge.ActionCarapaceBin(shell).Split()
			}
			return carapace.ActionValues()
		}),
		"show": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return just.ActionRecipes(rootCmd.Flag("justfile").Value.String())
		}),
		"tempdir":           carapace.ActionDirectories(),
		"usage":             carapace.ActionFiles(), // TODO verify
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
