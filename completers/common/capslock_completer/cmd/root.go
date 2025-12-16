package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/golang"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "capslock",
	Short: "Capslock is a capability analysis CLI for Go packages ",
	Long:  "https://github.com/google/capslock",
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

	rootCmd.Flags().StringS("buildtags", "buildtags", "", "command-separated list of build tags to use when loading packages")
	rootCmd.Flags().StringS("goarch", "goarch", "", "GOARCH value to use when loading packages")
	rootCmd.Flags().StringS("goos", "goos", "", "GOOS value to use when loading packages")
	rootCmd.Flags().BoolS("noisy", "noisy", false, "include output on unanalyzed function calls")
	rootCmd.Flags().StringS("output", "output", "", "output mode to use")
	rootCmd.Flags().StringS("packages", "packages", "", "target patterns to be analysed")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"buildtags": golang.ActionBuildTags(),
		"goarch":    golang.ActionArchitectures(),
		"goos":      golang.ActionOperatingSystems(),
		"output":    carapace.ActionValues("json", "m", "v", "graph", "compare"),
		"packages":  golang.ActionPackages(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if rootCmd.Flag("output").Value.String() == "compare" {
				return carapace.ActionFiles()
			}
			return carapace.ActionValues()
		}),
	)
}
