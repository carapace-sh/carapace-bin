package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/carapace-sh/carapace/pkg/traverse"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var rootCmd = &cobra.Command{
	Use:   "docker-compose",
	Short: "Docker Compose",
	Long:  "https://docs.docker.com/compose/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("all-resources", false, "Include all resources, even those not used by services")
	rootCmd.Flags().String("ansi", "", "Control when to print ANSI control characters (\"never\"|\"always\"|\"auto\")")
	rootCmd.Flags().Bool("compatibility", false, "Run compose in backward compatibility mode")
	rootCmd.PersistentFlags().Bool("dry-run", false, "Execute command in dry run mode")
	rootCmd.Flags().StringSlice("env-file", nil, "Specify an alternate environment file")
	rootCmd.Flags().StringSliceP("file", "f", nil, "Compose configuration files")
	rootCmd.Flags().Bool("no-ansi", false, "Do not print ANSI control characters (DEPRECATED)")
	rootCmd.Flags().String("parallel", "", "Control max parallelism, -1 for unlimited")
	rootCmd.Flags().StringSlice("profile", nil, "Specify a profile to enable")
	rootCmd.Flags().String("progress", "", "Set type of progress output (auto, tty, plain, json, quiet)")
	rootCmd.Flags().String("project-directory", "", "Specify an alternate working directory")
	rootCmd.Flags().StringP("project-name", "p", "", "Project name")
	rootCmd.Flags().Bool("verbose", false, "Show more output")
	rootCmd.Flags().BoolP("version", "v", false, "Show the Docker Compose version information")
	rootCmd.Flags().String("workdir", "", "DEPRECATED! USE --project-directory INSTEAD.")
	rootCmd.Flag("no-ansi").Hidden = true
	rootCmd.Flag("verbose").Hidden = true
	rootCmd.Flag("version").Hidden = true
	rootCmd.Flag("workdir").Hidden = true

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"ansi":              carapace.ActionValues("auto", "never", "always").StyleF(style.ForKeyword),
		"env-file":          carapace.ActionFiles(),
		"file":              carapace.ActionFiles(),
		"progress":          carapace.ActionValues("auto", "tty", "plain", "quiet"),
		"project-directory": carapace.ActionDirectories(),
	})

	carapace.Gen(rootCmd).PreInvoke(func(cmd *cobra.Command, flag *pflag.Flag, action carapace.Action) carapace.Action {
		return action.ChdirF(traverse.Flag(rootCmd.Flag("project-directory")))
	})
}
