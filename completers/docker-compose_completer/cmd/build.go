package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/docker-compose_completer/cmd/action"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:   "build [OPTIONS] [SERVICE...]",
	Short: "Build or rebuild services",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(buildCmd).Standalone()

	buildCmd.Flags().StringArray("build-arg", []string{}, "Set build-time variables for services.")
	buildCmd.Flags().Bool("compress", true, "Compress the build context using gzip. DEPRECATED")
	buildCmd.Flags().Bool("force-rm", true, "Always remove intermediate containers. DEPRECATED")
	buildCmd.Flags().StringP("memory", "m", "", "Set memory limit for the build container. Not supported on buildkit yet.")
	buildCmd.Flags().Bool("no-cache", false, "Do not use cache when building the image")
	buildCmd.Flags().Bool("no-rm", false, "Do not remove intermediate containers after a successful build. DEPRECATED")
	buildCmd.Flags().Bool("parallel", true, "Build images in parallel. DEPRECATED")
	buildCmd.Flags().String("progress", "auto", "Set type of progress output (auto, tty, plain, quiet)")
	buildCmd.Flags().Bool("pull", false, "Always attempt to pull a newer version of the image.")
	buildCmd.Flags().Bool("push", false, "Push service images.")
	buildCmd.Flags().BoolP("quiet", "q", false, "Don't print anything to STDOUT")
	buildCmd.Flags().String("ssh", "", "Set SSH authentications used when building service images. (use 'default' for using your default SSH Agent)")
	rootCmd.AddCommand(buildCmd)

	carapace.Gen(buildCmd).FlagCompletion(carapace.ActionMap{
		"progress": carapace.ActionValues("auto", "tty", "plain", "quiet"),
		"pull":     carapace.ActionValues("always", "missing", "never").StyleF(style.ForKeyword),
	})

	carapace.Gen(buildCmd).PositionalAnyCompletion(
		action.ActionServices(buildCmd).FilterArgs(),
	)
}
