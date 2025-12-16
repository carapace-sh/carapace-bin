package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/helix"
	"github.com/carapace-sh/carapace/pkg/traverse"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var rootCmd = &cobra.Command{
	Use:   "helix",
	Short: "A post-modern text editor",
	Long:  "https://helix-editor.com/",
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

	rootCmd.Flags().StringP("config", "c", "", "Specifies a file to use for configuration")
	rootCmd.Flags().StringP("grammar", "g", "", "Fetches or builds tree-sitter grammars listed in languages.toml")
	rootCmd.Flags().String("health", "", "Checks for potential errors in editor setup")
	rootCmd.Flags().BoolP("help", "h", false, "Prints help information")
	rootCmd.Flags().Bool("hsplit", false, "Splits all given files horizontally into different windows")
	rootCmd.Flags().String("log", "", "Specifies a file to use for logging")
	rootCmd.Flags().Bool("tutor", false, "Loads the tutorial")
	rootCmd.Flags().CountS("v", "v", "Increases logging verbosity each use for up to 3 times")
	rootCmd.Flags().BoolP("version", "V", false, "Prints version information")
	rootCmd.Flags().Bool("vsplit", false, "Splits all given files vertically into different windows")
	rootCmd.Flags().StringP("working-dir", "w", "", "Specify an initial working directory")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"config":      carapace.ActionFiles(),
		"grammar":     carapace.ActionValues("fetch", "build"),
		"health":      helix.ActionLanguages(),
		"log":         carapace.ActionFiles(),
		"working-dir": carapace.ActionDirectories(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)

	carapace.Gen(rootCmd).PreInvoke(func(cmd *cobra.Command, flag *pflag.Flag, action carapace.Action) carapace.Action {
		return action.ChdirF(traverse.Flag(rootCmd.Flag("working-dir")))
	})
}
