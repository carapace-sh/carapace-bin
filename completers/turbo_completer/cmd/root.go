package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var rootCmd = &cobra.Command{
	Use:   "turbo",
	Short: "The build system that makes ship happen",
	Long:  "https://turbo.build/repo",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.PersistentFlags().String("api", "", "Override the endpoint for API calls")
	rootCmd.PersistentFlags().Bool("color", false, "Force color usage in the terminal")
	rootCmd.PersistentFlags().String("cpuprofile", "", "Specify a file to save a cpu profile")
	rootCmd.PersistentFlags().String("cwd", "", "The directory in which to run turbo")
	rootCmd.PersistentFlags().String("heap", "", "Specify a file to save a pprof heap profile")
	rootCmd.PersistentFlags().BoolP("help", "h", false, "Print help")
	rootCmd.PersistentFlags().String("login", "", "Override the login endpoint")
	rootCmd.PersistentFlags().Bool("no-color", false, "Suppress color usage in the terminal")
	rootCmd.PersistentFlags().Bool("preflight", false, "When enabled, turbo will precede HTTP requests with an OPTIONS request for authorization")
	rootCmd.PersistentFlags().Bool("single-package", false, "Run turbo in single-package mode")
	rootCmd.PersistentFlags().Bool("skip-infer", false, "Skip any attempts to infer which version of Turbo the project is configured to use")
	rootCmd.PersistentFlags().String("team", "", "Set the team slug for API calls")
	rootCmd.PersistentFlags().String("token", "", "Set the auth token for API calls")
	rootCmd.PersistentFlags().String("trace", "", "Specify a file to save a pprof trace")
	rootCmd.PersistentFlags().String("verbosity", "", "Verbosity level")
	rootCmd.PersistentFlags().Bool("version", false, "")
	addRunFlags(rootCmd)

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"cpuprofile": carapace.ActionFiles(),
		"cwd":        carapace.ActionDirectories(),
		"heap":       carapace.ActionFiles(),
		"trace":      carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PreInvoke(func(cmd *cobra.Command, flag *pflag.Flag, action carapace.Action) carapace.Action {
		return action.Chdir(rootCmd.Flag("cwd").Value.String())
	})

}
