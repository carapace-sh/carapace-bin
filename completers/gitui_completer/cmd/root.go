package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gitui",
	Short: "blazing fast terminal-ui for git",
	Long:  "https://github.com/extrawurst/gitui",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().String("bugreport", "", "Generate a bug report")
	rootCmd.Flags().StringP("directory", "d", "", "Set the git directory")
	rootCmd.Flags().BoolP("help", "h", false, "Print help information")
	rootCmd.Flags().BoolP("logging", "l", false, "Stores logging output into a cache directory")
	rootCmd.Flags().Bool("polling", false, "Poll folder for changes instead of using file system events")
	rootCmd.Flags().StringP("theme", "t", "", "Set the color theme")
	rootCmd.Flags().BoolP("version", "V", false, "Print version information")
	rootCmd.Flags().StringP("workdir", "w", "", "Set the working directory")

	// TODO theme completion
	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"directory": carapace.ActionDirectories(),
		"workdir":   carapace.ActionDirectories(),
	})
}
