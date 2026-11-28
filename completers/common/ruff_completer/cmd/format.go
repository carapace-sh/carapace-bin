package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var formatCmd = &cobra.Command{
	Use:   "format",
	Short: "Run the Ruff formatter on the given files or directories",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(formatCmd).Standalone()

	formatCmd.Flags().String("cache-dir", "", "Path to the cache directory")
	formatCmd.Flags().Bool("check", false, "Avoid writing any formatted files back; instead, exit with a non-zero status code if any files would have been modified, and zero otherwise")
	formatCmd.Flags().Bool("diff", false, "Avoid writing any formatted files back; instead, exit with a non-zero status code and the difference between the current file and how the formatted file would look like")
	formatCmd.Flags().StringSlice("exclude", nil, "List of paths, used to omit files and/or directories from analysis")
	formatCmd.Flags().Bool("exit-non-zero-on-fix", false, "Exit with a non-zero status code if any files were modified via format, even if all files were formatted successfully")
	formatCmd.Flags().Bool("exit-non-zero-on-format", false, "Exit with a non-zero status code if any files were modified via format, even if all files were formatted successfully")
	formatCmd.Flags().StringSlice("extend-exclude", nil, "Like --exclude, but adds additional files and directories on top of those already excluded")
	formatCmd.Flags().StringSlice("extension", nil, "List of mappings from file extension to language (one of `python`, `ipynb`, `pyi`). For example, to treat `.ipy` files as IPython notebooks, use `--extension ipy:ipynb`")
	formatCmd.Flags().Bool("force-exclude", false, "Enforce exclusions, even for paths passed to Ruff directly on the command-line. Use `--no-force-exclude` to disable")
	formatCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	formatCmd.Flags().String("line-length", "", "Set the line-length")
	formatCmd.Flags().BoolP("no-cache", "n", false, "Disable cache reads")
	formatCmd.Flags().Bool("no-force-exclude", false, "")
	formatCmd.Flags().Bool("no-preview", false, "")
	formatCmd.Flags().Bool("no-respect-gitignore", false, "")
	formatCmd.Flags().String("output-format", "", "Output serialization format for violations, when used with `--check`. The default serialization format is \"full\"")
	formatCmd.Flags().Bool("preview", false, "Enable preview mode; enables unstable formatting. Use `--no-preview` to disable")
	formatCmd.Flags().String("range", "", "When specified, Ruff will try to only format the code in the given range.")
	formatCmd.Flags().Bool("respect-gitignore", false, "Respect file exclusions via `.gitignore` and other standard ignore files. Use `--no-respect-gitignore` to disable")
	formatCmd.Flags().String("stdin-filename", "", "The name of the file when passing it through stdin")
	formatCmd.Flags().String("target-version", "", "The minimum Python version that should be supported")
	formatCmd.Flag("exit-non-zero-on-fix").Hidden = true
	formatCmd.Flag("no-force-exclude").Hidden = true
	formatCmd.Flag("no-preview").Hidden = true
	formatCmd.Flag("no-respect-gitignore").Hidden = true
	rootCmd.AddCommand(formatCmd)
}
