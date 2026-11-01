package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-tool-dist",
	Short: "Dist helps bootstrap, build, and test the Go distribution",
	Long:  "https://pkg.go.dev/cmd/dist",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("v", "v", false, "verbosity")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"banner", "print installation banner",
			"bootstrap", "rebuild everything",
			"clean", "deletes all built files",
			"env", "print environment",
			"install", "install individual directory",
			"list", "list all supported platforms",
			"test", "run Go test(s)",
			"version", "print Go version",
		),
	)
}
