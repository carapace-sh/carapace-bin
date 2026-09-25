package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mise",
	Short: "The front-end to your dev env",
	Long:  "https://mise.jdx.dev",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.PersistentFlags().StringP("cd", "C", "", "Change directory before running command")
	rootCmd.PersistentFlags().StringP("env", "E", "", "Set the environment for loading `mise.<ENV>.toml`")
	rootCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.PersistentFlags().StringP("jobs", "j", "", "How many jobs to run in parallel")
	rootCmd.PersistentFlags().Bool("locked", false, "Require lockfile URLs to be present during installation")
	rootCmd.PersistentFlags().BoolP("quiet", "q", false, "Suppress non-error messages")
	rootCmd.PersistentFlags().Bool("raw", false, "Read/write directly to stdin/stdout/stderr instead of by line")
	rootCmd.PersistentFlags().Bool("silent", false, "Suppress all task output and mise non-error messages")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Show extra output")
	rootCmd.Flags().BoolP("version", "V", false, "Print version")
	rootCmd.PersistentFlags().BoolP("yes", "y", false, "Answer yes to all confirmation prompts")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"cd":  carapace.ActionDirectories(),
		"env": carapace.ActionValues("development", "production", "staging", "test"),
	})
}
