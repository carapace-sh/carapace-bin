package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var analyze_graphCmd = &cobra.Command{
	Use:   "graph",
	Short: "Generate a map of Python file dependencies or dependents",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(analyze_graphCmd).Standalone()

	analyze_graphCmd.Flags().Bool("detect-string-imports", false, "Attempt to detect imports from string literals")
	analyze_graphCmd.Flags().String("direction", "dependencies", "The direction of the import map. By default, generates a dependency map, i.e., a map from file to files that it depends on. Use `--direction dependents` to generate a map from file to files that depend on it")
	analyze_graphCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	analyze_graphCmd.Flags().String("min-dots", "", "The minimum number of dots in a string import to consider it a valid import")
	analyze_graphCmd.Flags().Bool("no-preview", false, "")
	analyze_graphCmd.Flags().Bool("no-type-checking-imports", false, "")
	analyze_graphCmd.Flags().Bool("preview", false, "Enable preview mode. Use `--no-preview` to disable")
	analyze_graphCmd.Flags().String("python", "", "Path to a virtual environment to use for resolving additional dependencies")
	analyze_graphCmd.Flags().String("target-version", "", "The minimum Python version that should be supported")
	analyze_graphCmd.Flags().Bool("type-checking-imports", false, "Include imports that are only used for type checking (i.e., imports within `if TYPE_CHECKING:` blocks). Use `--no-type-checking-imports` to exclude imports that are only used for type checking")
	analyze_graphCmd.Flag("no-preview").Hidden = true
	analyze_graphCmd.Flag("no-type-checking-imports").Hidden = true
	analyzeCmd.AddCommand(analyze_graphCmd)
}
