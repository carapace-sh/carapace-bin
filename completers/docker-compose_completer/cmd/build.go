package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build or rebuild services",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	buildCmd.Flags().String("build-arg", "", "Set build-time variables for services.")
	buildCmd.Flags().Bool("compress", false, "Compress the build context using gzip.")
	buildCmd.Flags().Bool("force-rm", false, "Always remove intermediate containers.")
	buildCmd.Flags().StringP("memory", "m", "", "Set memory limit for the build container.")
	buildCmd.Flags().Bool("no-cache", false, "Do not use cache when building the image.")
	buildCmd.Flags().Bool("no-rm", false, "Do not remove intermediate containers after a successful build.")
	buildCmd.Flags().Bool("parallel", false, "Build images in parallel.")
	buildCmd.Flags().String("progress", "", "Set type of progress output (auto, plain, tty).")
	buildCmd.Flags().Bool("pull", false, "Always attempt to pull a newer version of the image.")
	buildCmd.Flags().BoolP("quiet", "q", false, "Don't print anything to STDOUT")
	rootCmd.AddCommand(buildCmd)

	carapace.Gen(buildCmd).FlagCompletion(carapace.ActionMap{
		"progress": carapace.ActionValues("auto", "plain", "tty"),
	})

	carapace.Gen(buildCmd).PositionalAnyCompletion(ActionServices())
}
