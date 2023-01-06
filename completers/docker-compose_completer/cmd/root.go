package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "docker-compose",
	Short: "Define and run multi-container applications with Docker",
	Long:  "https://docs.docker.com/compose/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()
	rootCmd.Flags().String("ansi", "auto", "Control when to print ANSI control characters (\"never\"|\"always\"|\"auto\")")
	rootCmd.Flags().Bool("compatibility", false, "Run compose in backward compatibility mode")
	rootCmd.Flags().String("env-file", "", "Specify an alternate environment file.")
	rootCmd.Flags().StringArrayP("file", "f", []string{}, "Compose configuration files")
	rootCmd.Flags().Bool("no-ansi", false, "Do not print ANSI control characters (DEPRECATED)")
	rootCmd.Flags().Int("parallel", -1, "Control max parallelism, -1 for unlimited")
	rootCmd.Flags().StringArray("profile", []string{}, "Specify a profile to enable")
	rootCmd.Flags().String("project-directory", "", "Specify an alternate working directory")
	rootCmd.Flags().StringP("project-name", "p", "", "Project name")
	rootCmd.Flags().Bool("verbose", false, "Show more output")
	rootCmd.Flags().BoolP("version", "v", false, "Show the Docker Compose version information")
	rootCmd.Flags().String("workdir", "", "DEPRECATED! USE --project-directory INSTEAD.")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"ansi":              carapace.ActionValues("auto", "never", "always").StyleF(style.ForKeyword),
		"env-file":          carapace.ActionFiles(),
		"file":              carapace.ActionFiles(),
		"project-directory": carapace.ActionDirectories(),
	})
}
