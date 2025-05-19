package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:     "build [OPTIONS] PATH | URL | -",
	Short:   "Start a build",
	Aliases: []string{"b"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(buildCmd).Standalone()
	buildCmd.Flags().StringSlice("add-host", nil, "Add a custom host-to-IP mapping (format: \"host:ip\")")
	buildCmd.Flags().StringSlice("allow", nil, "Allow extra privileged entitlement (e.g., \"network.host\", \"security.insecure\")")
	buildCmd.Flags().StringArray("build-arg", nil, "Set build-time variables")
	buildCmd.Flags().StringArray("build-context", nil, "Additional build contexts (e.g., name=path)")
	buildCmd.Flags().StringArray("cache-from", nil, "External cache sources (e.g., \"user/app:cache\", \"type=local,src=path/to/dir\")")
	buildCmd.Flags().StringArray("cache-to", nil, "Cache export destinations (e.g., \"user/app:cache\", \"type=local,dest=path/to/dir\")")
	buildCmd.Flags().String("cgroup-parent", "", "Optional parent cgroup for the container")
	buildCmd.Flags().Bool("compress", false, "Compress the build context using gzip")
	buildCmd.Flags().Int64("cpu-period", 0, "Limit the CPU CFS (Completely Fair Scheduler) period")
	buildCmd.Flags().Int64("cpu-quota", 0, "Limit the CPU CFS (Completely Fair Scheduler) quota")
	buildCmd.Flags().Int64P("cpu-shares", "c", 0, "CPU shares (relative weight)")
	buildCmd.Flags().String("cpuset-cpus", "", "CPUs in which to allow execution (\"0-3\", \"0,1\")")
	buildCmd.Flags().String("cpuset-mems", "", "MEMs in which to allow execution (\"0-3\", \"0,1\")")
	buildCmd.Flags().StringP("file", "f", "", "Name of the Dockerfile (default: \"PATH/Dockerfile\")")
	buildCmd.Flags().Bool("force-rm", false, "Always remove intermediate containers")
	buildCmd.Flags().String("iidfile", "", "Write the image ID to the file")
	buildCmd.Flags().String("isolation", "", "Container isolation technology")
	buildCmd.Flags().StringArray("label", nil, "Set metadata for an image")
	buildCmd.Flags().Bool("load", false, "Shorthand for \"--output=type=docker\"")
	buildCmd.Flags().StringP("memory", "m", "", "Memory limit")
	buildCmd.Flags().String("memory-swap", "", "Swap limit equal to memory plus swap: \"-1\" to enable unlimited swap")
	buildCmd.Flags().String("metadata-file", "", "Write build result metadata to the file")
	buildCmd.Flags().String("network", "default", "Set the networking mode for the \"RUN\" instructions during build")
	buildCmd.Flags().Bool("no-cache", false, "Do not use cache when building the image")
	buildCmd.Flags().StringArray("no-cache-filter", nil, "Do not cache specified stages")
	buildCmd.Flags().StringArrayP("output", "o", nil, "Output destination (format: \"type=local,dest=path\")")
	buildCmd.Flags().StringArray("platform", nil, "Set target platform for build")
	buildCmd.Flags().String("progress", "auto", "Set type of progress output (\"auto\", \"plain\", \"tty\"). Use plain to show container output")
	buildCmd.Flags().Bool("pull", false, "Always attempt to pull all referenced images")
	buildCmd.Flags().Bool("push", false, "Shorthand for \"--output=type=registry\"")
	buildCmd.Flags().BoolP("quiet", "q", false, "Suppress the build output and print image ID on success")
	buildCmd.Flags().Bool("rm", false, "Remove intermediate containers after a successful build")
	buildCmd.Flags().StringArray("secret", nil, "Secret to expose to the build (format: \"id=mysecret[,src=/local/secret]\")")
	buildCmd.Flags().StringSlice("security-opt", nil, "Security options")
	buildCmd.Flags().String("shm-size", "", "Size of \"/dev/shm\"")
	buildCmd.Flags().Bool("squash", false, "Squash newly built layers into a single new layer")
	buildCmd.Flags().StringArray("ssh", nil, "SSH agent socket or keys to expose to the build (format: \"default|<id>[=<socket>|<key>[,<key>]]\")")
	buildCmd.Flags().StringArrayP("tag", "t", nil, "Name and optionally a tag (format: \"name:tag\")")
	buildCmd.Flags().String("target", "", "Set the target build stage to build")
	buildCmd.Flags().String("ulimit", "", "Ulimit options")
	rootCmd.AddCommand(buildCmd)

	// TODO flag completion
	carapace.Gen(buildCmd).FlagCompletion(carapace.ActionMap{
		"cgroup-parent": os.ActionCgroups(),
		"file":          carapace.ActionFiles(),
		"iidfile":       carapace.ActionFiles(),
		"metadata-file": carapace.ActionFiles(),
		"progress":      carapace.ActionValues("auto", "plain", "tty"),
	})

	carapace.Gen(buildCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
