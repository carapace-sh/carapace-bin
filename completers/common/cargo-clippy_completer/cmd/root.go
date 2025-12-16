package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/cargo/clippy"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cargo-clippy",
	Short: "Checks a package to catch common mistakes and improve your Rust code",
	Long:  "https://github.com/arcnmx/cargo-clippy",
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

	rootCmd.Flags().Bool("fix", false, "Automatically apply lint suggestions. This flag implies `--no-deps`")
	rootCmd.Flags().BoolP("help", "h", false, "Print this message")
	rootCmd.Flags().Bool("no-deps", false, "Run Clippy only on the given crate, without linting the dependencies")
	rootCmd.Flags().BoolP("version", "V", false, "Print version info and exit")

	carapace.Gen(rootCmd).DashAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			dashCmd := &cobra.Command{}
			carapace.Gen(dashCmd).Standalone()

			dashCmd.Flags().StringArrayP("allow", "A", nil, "Set lint allowed")
			dashCmd.Flags().StringArrayP("deny", "D", nil, "Set lint denied")
			dashCmd.Flags().StringArrayP("forbid", "F", nil, "Set lint forbidden")
			dashCmd.Flags().StringArrayP("warn", "W", nil, "Set lint warnings")

			carapace.Gen(dashCmd).FlagCompletion(carapace.ActionMap{
				"allow":  actionLintsAndCategories(),
				"deny":   actionLintsAndCategories(),
				"forbid": actionLintsAndCategories(),
				"warn":   actionLintsAndCategories(),
			})

			return carapace.ActionExecute(dashCmd)
		}),
	)
}

func actionLintsAndCategories() carapace.Action {
	return carapace.Batch(
		clippy.ActionCategories(),
		clippy.ActionLints(),
	).ToA()
}
