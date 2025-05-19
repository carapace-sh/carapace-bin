package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "bloop",
	Short: "Build and test Scala code",
	Long:  "https://scalacenter.github.io/bloop/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.PersistentFlags().StringP("config-dir", "c", "", "file path to the bloop config directory")
	rootCmd.PersistentFlags().String("debug", "", "Debug the execution of a concrete task")
	rootCmd.PersistentFlags().Bool("no-color", false, "do not color output")
	rootCmd.PersistentFlags().Bool("verbose", false, "print out debugging information to stderr")
	rootCmd.PersistentFlags().BoolP("version", "v", false, "print the root section at the beginning of the execution")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"config-dir": carapace.ActionDirectories(),
		"debug":      carapace.ActionValues("all", "file-watching", "compilation", "test", "bsp", "link"),
	})

}
