package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/docker-compose_completer/cmd/action"
	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:   "build [OPTIONS] [SERVICE...]",
	Short: "Build or rebuild services",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(buildCmd).Standalone()

	buildCmd.Flags().StringSlice("build-arg", []string{}, "Set build-time variables for services")
	buildCmd.Flags().String("builder", "", "Set builder to use")
	buildCmd.Flags().Bool("compress", false, "Compress the build context using gzip. DEPRECATED")
	buildCmd.Flags().Bool("force-rm", false, "Always remove intermediate containers. DEPRECATED")
	buildCmd.Flags().StringP("memory", "m", "", "Set memory limit for the build container. Not supported by BuildKit.")
	buildCmd.Flags().Bool("no-cache", false, "Do not use cache when building the image")
	buildCmd.Flags().Bool("no-rm", false, "Do not remove intermediate containers after a successful build. DEPRECATED")
	buildCmd.Flags().Bool("parallel", false, "Build images in parallel. DEPRECATED")
	buildCmd.Flags().String("progress", "", "Set type of ui output (auto, tty, plain, json, quiet)")
	buildCmd.Flags().Bool("pull", false, "Always attempt to pull a newer version of the image")
	buildCmd.Flags().Bool("push", false, "Push service images")
	buildCmd.Flags().BoolP("quiet", "q", false, "Don't print anything to STDOUT")
	buildCmd.Flags().String("ssh", "", "Set SSH authentications used when building service images. (use 'default' for using your default SSH Agent)")
	buildCmd.Flags().Bool("with-dependencies", false, "Also build dependencies (transitively)")
	buildCmd.Flag("compress").Hidden = true
	buildCmd.Flag("force-rm").Hidden = true
	buildCmd.Flag("no-rm").Hidden = true
	buildCmd.Flag("parallel").Hidden = true
	buildCmd.Flag("progress").Hidden = true
	rootCmd.AddCommand(buildCmd)

	// TODO builder
	carapace.Gen(buildCmd).FlagCompletion(carapace.ActionMap{
		"progress": carapace.ActionValues("auto", "tty", "plain", "json", "quiet"),
	})

	carapace.Gen(buildCmd).PositionalAnyCompletion(
		action.ActionServices(buildCmd).FilterArgs(),
	)
}
