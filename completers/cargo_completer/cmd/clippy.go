package cmd

import (
	"os"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/cargo_completer/cmd/action"
	"github.com/spf13/cobra"
)

var clippyCmd = &cobra.Command{
	Use:   "clippy",
	Short: "Checks a package to catch common mistakes and improve your Rust code",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clippyCmd).Standalone()

	clippyCmd.Flags().Bool("fix", false, "Automatically apply lint suggestions. This flag implies `--no-deps`")
	clippyCmd.Flags().BoolP("help", "h", false, "Print this message")
	clippyCmd.Flags().Bool("no-deps", false, "Run Clippy only on the given crate, without linting the dependencies")
	clippyCmd.Flags().BoolP("version", "V", false, "Print version info and exit")

	for _, arg := range os.Args {
		if arg == "--" {
			clippyCmd.Flags().StringArrayP("allow", "A", []string{}, "Set lint allowed")
			clippyCmd.Flags().StringArrayP("deny", "D", []string{}, "Set lint denied")
			clippyCmd.Flags().StringArrayP("forbid", "F", []string{}, "Set lint forbidden")
			clippyCmd.Flags().StringArrayP("warn", "W", []string{}, "Set lint warnings")

			carapace.Gen(clippyCmd).FlagCompletion(carapace.ActionMap{
				"allow":  action.ActionLints(),
				"deny":   action.ActionLints(),
				"forbid": action.ActionLints(),
				"warn":   action.ActionLints(),
			})

			break
		}
	}

	rootCmd.AddCommand(clippyCmd)
}
