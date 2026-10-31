package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Run Ruff on the given files or directories",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(checkCmd).Standalone()

	checkCmd.Flags().String("add-ignore", "", "Enable automatic additions of `ruff: ignore` comments to failing lines. Optionally provide a reason to append after the codes. In preview, add suppression comments with rule names instead")
	checkCmd.Flags().String("add-noqa", "", "Enable automatic additions of `noqa` directives to failing lines. Optionally provide a reason to append after the codes")
	checkCmd.Flags().String("cache-dir", "", "Path to the cache directory")
	checkCmd.Flags().Bool("diff", false, "Avoid writing any fixed files back; instead, output a diff for each changed file to stdout, and exit 0 if there are no diffs. Implies `--fix-only`")
	checkCmd.Flags().String("dummy-variable-rgx", "", "Regular expression matching the name of dummy variables")
	checkCmd.Flags().StringSlice("exclude", nil, "List of paths, used to omit files and/or directories from analysis")
	checkCmd.Flags().Bool("exit-non-zero-on-fix", false, "Exit with a non-zero status code if any files were modified via fix, even if no lint violations remain")
	checkCmd.Flags().BoolP("exit-zero", "e", false, "Exit with status code \"0\", even upon detecting lint violations")
	checkCmd.Flags().StringSlice("extend-exclude", nil, "Like --exclude, but adds additional files and directories on top of those already excluded")
	checkCmd.Flags().StringSlice("extend-fixable", nil, "Like --fixable, but adds additional rule codes on top of those already specified")
	checkCmd.Flags().StringSlice("extend-ignore", nil, "Like --ignore. (Deprecated: You can just use --ignore instead.)")
	checkCmd.Flags().StringSlice("extend-per-file-ignores", nil, "Like `--per-file-ignores`, but adds additional ignores on top of those already specified")
	checkCmd.Flags().StringSlice("extend-select", nil, "Like --select, but adds additional rule codes on top of those already specified")
	checkCmd.Flags().StringSlice("extend-unfixable", nil, "Like --unfixable. (Deprecated: You can just use --unfixable instead.)")
	checkCmd.Flags().StringSlice("extension", nil, "List of mappings from file extension to language (one of `python`, `ipynb`, `pyi`). For example, to treat `.ipy` files as IPython notebooks, use `--extension ipy:ipynb`")
	checkCmd.Flags().Bool("fix", false, "Apply fixes to resolve lint violations. Use `--no-fix` to disable or `--unsafe-fixes` to include unsafe fixes")
	checkCmd.Flags().Bool("fix-only", false, "Apply fixes to resolve lint violations, but don't report on, or exit non-zero for, leftover violations. Implies `--fix`. Use `--no-fix-only` to disable or `--unsafe-fixes` to include unsafe fixes")
	checkCmd.Flags().StringSlice("fixable", nil, "List of rule codes to treat as eligible for fix. Only applicable when fix itself is enabled (e.g., via `--fix`)")
	checkCmd.Flags().Bool("force-exclude", false, "Enforce exclusions, even for paths passed to Ruff directly on the command-line. Use `--no-force-exclude` to disable")
	checkCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	checkCmd.Flags().StringSlice("ignore", nil, "Comma-separated list of rule codes to disable")
	checkCmd.Flags().Bool("ignore-noqa", false, "Ignore any `# noqa` comments")
	checkCmd.Flags().String("line-length", "", "Set the line-length for length-associated rules and automatic formatting")
	checkCmd.Flags().BoolP("no-cache", "n", false, "Disable cache reads")
	checkCmd.Flags().Bool("no-fix", false, "")
	checkCmd.Flags().Bool("no-fix-only", false, "")
	checkCmd.Flags().Bool("no-force-exclude", false, "")
	checkCmd.Flags().Bool("no-preview", false, "")
	checkCmd.Flags().Bool("no-respect-gitignore", false, "")
	checkCmd.Flags().Bool("no-show-fixes", false, "")
	checkCmd.Flags().Bool("no-unsafe-fixes", false, "")
	checkCmd.Flags().StringP("output-file", "o", "", "Specify file to write the linter output to (default: stdout)")
	checkCmd.Flags().String("output-format", "", "Output serialization format for violations. The default serialization format is \"full\"")
	checkCmd.Flags().StringSlice("per-file-ignores", nil, "List of mappings from file pattern to code to exclude")
	checkCmd.Flags().Bool("preview", false, "Enable preview mode; checks will include unstable rules and fixes. Use `--no-preview` to disable")
	checkCmd.Flags().Bool("respect-gitignore", false, "Respect file exclusions via `.gitignore` and other standard ignore files. Use `--no-respect-gitignore` to disable")
	checkCmd.Flags().StringSlice("select", nil, "Comma-separated list of rule codes to enable (or ALL, to enable all rules)")
	checkCmd.Flags().Bool("show-files", false, "See the files Ruff will be run against with the current settings")
	checkCmd.Flags().Bool("show-fixes", false, "Show an enumeration of all fixed lint violations. Use `--no-show-fixes` to disable")
	checkCmd.Flags().Bool("show-settings", false, "See the settings Ruff will use to lint a given Python file")
	checkCmd.Flags().Bool("statistics", false, "Show counts for every rule with at least one violation")
	checkCmd.Flags().String("stdin-filename", "", "The name of the file when passing it through stdin")
	checkCmd.Flags().String("target-version", "", "The minimum Python version that should be supported")
	checkCmd.Flags().StringSlice("unfixable", nil, "List of rule codes to treat as ineligible for fix. Only applicable when fix itself is enabled (e.g., via `--fix`)")
	checkCmd.Flags().Bool("unsafe-fixes", false, "Include fixes that may not retain the original intent of the code. Use `--no-unsafe-fixes` to disable")
	checkCmd.Flags().BoolP("watch", "w", false, "Run in watch mode by re-running whenever files change")
	checkCmd.Flag("add-ignore").NoOptDefVal = " "
	checkCmd.Flag("add-noqa").NoOptDefVal = " "
	checkCmd.Flag("dummy-variable-rgx").Hidden = true
	checkCmd.Flag("extend-ignore").Hidden = true
	checkCmd.Flag("extend-unfixable").Hidden = true
	checkCmd.Flag("line-length").Hidden = true
	checkCmd.Flag("no-fix").Hidden = true
	checkCmd.Flag("no-fix-only").Hidden = true
	checkCmd.Flag("no-force-exclude").Hidden = true
	checkCmd.Flag("no-preview").Hidden = true
	checkCmd.Flag("no-respect-gitignore").Hidden = true
	checkCmd.Flag("no-show-fixes").Hidden = true
	checkCmd.Flag("no-unsafe-fixes").Hidden = true
	rootCmd.AddCommand(checkCmd)
}
