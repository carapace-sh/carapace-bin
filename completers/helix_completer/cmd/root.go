package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/helix"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "helix",
	Short: "A post-modern text editor",
	Long:  "https://helix-editor.com/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func ExecuteHx() error {
	rootCmd.Use = "hx"
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

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"config":  carapace.ActionFiles(),
		"grammar": carapace.ActionValues("fetch", "build"),
		"health":  helix.ActionLanguages(),
		"log":     carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
