package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var builder_buildCmd = &cobra.Command{
	Use:   "build [OPTIONS] PATH | URL | -",
	Short: "Build an image from a Dockerfile",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(builder_buildCmd).Standalone()

	builder_buildCmd.Flags().String("add-host", "", "Add a custom host-to-IP mapping (\"host:ip\")")
	builder_buildCmd.Flags().String("build-arg", "", "Set build-time variables")
	builder_buildCmd.Flags().StringSlice("cache-from", nil, "Images to consider as cache sources")
	builder_buildCmd.Flags().String("cgroup-parent", "", "Set the parent cgroup for the \"RUN\" instructions during build")
	builder_buildCmd.Flags().Bool("compress", false, "Compress the build context using gzip")
	builder_buildCmd.Flags().String("cpu-period", "", "Limit the CPU CFS (Completely Fair Scheduler) period")
	builder_buildCmd.Flags().String("cpu-quota", "", "Limit the CPU CFS (Completely Fair Scheduler) quota")
	builder_buildCmd.Flags().StringP("cpu-shares", "c", "", "CPU shares (relative weight)")
	builder_buildCmd.Flags().String("cpuset-cpus", "", "CPUs in which to allow execution (0-3, 0,1)")
	builder_buildCmd.Flags().String("cpuset-mems", "", "MEMs in which to allow execution (0-3, 0,1)")
	builder_buildCmd.Flags().Bool("disable-content-trust", false, "Skip image verification")
	builder_buildCmd.Flags().StringP("file", "f", "", "Name of the Dockerfile (Default is \"PATH/Dockerfile\")")
	builder_buildCmd.Flags().Bool("force-rm", false, "Always remove intermediate containers")
	builder_buildCmd.Flags().String("iidfile", "", "Write the image ID to the file")
	builder_buildCmd.Flags().String("isolation", "", "Container isolation technology")
	builder_buildCmd.Flags().String("label", "", "Set metadata for an image")
	builder_buildCmd.Flags().StringP("memory", "m", "", "Memory limit")
	builder_buildCmd.Flags().String("memory-swap", "", "Swap limit equal to memory plus swap: -1 to enable unlimited swap")
	builder_buildCmd.Flags().String("network", "", "Set the networking mode for the RUN instructions during build")
	builder_buildCmd.Flags().Bool("no-cache", false, "Do not use cache when building the image")
	builder_buildCmd.Flags().String("platform", "", "Set platform if server is multi-platform capable")
	builder_buildCmd.Flags().Bool("pull", false, "Always attempt to pull a newer version of the image")
	builder_buildCmd.Flags().BoolP("quiet", "q", false, "Suppress the build output and print image ID on success")
	builder_buildCmd.Flags().Bool("rm", false, "Remove intermediate containers after a successful build")
	builder_buildCmd.Flags().StringSlice("security-opt", nil, "Security options")
	builder_buildCmd.Flags().String("shm-size", "", "Size of \"/dev/shm\"")
	builder_buildCmd.Flags().Bool("squash", false, "Squash newly built layers into a single new layer")
	builder_buildCmd.Flags().StringP("tag", "t", "", "Name and optionally a tag in the \"name:tag\" format")
	builder_buildCmd.Flags().String("target", "", "Set the target build stage to build.")
	builder_buildCmd.Flags().String("ulimit", "", "Ulimit options")
	builderCmd.AddCommand(builder_buildCmd)

	// TODO flag completion
	carapace.Gen(builder_buildCmd).FlagCompletion(carapace.ActionMap{
		"add-host":      carapace.ActionValues(),
		"cache-from":    docker.ActionRepositoryTags(), // TODO verify
		"cgroup-parent": os.ActionCgroups(),            // TODO verify
		"file":          carapace.ActionFiles(),
		"iidfile":       carapace.ActionFiles(),
		"isolation":     carapace.ActionValues(),
		"network":       carapace.ActionValues(),
		"platform":      carapace.ActionValues(),
		"security-opt":  carapace.ActionValues(),
		"target":        carapace.ActionValues(),
	})

	carapace.Gen(builder_buildCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
