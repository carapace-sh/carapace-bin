package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var farm_buildCmd = &cobra.Command{
	Use:   "build [options] [CONTEXT]",
	Short: "Build a container image for multiple architectures",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(farm_buildCmd).Standalone()

	farm_buildCmd.Flags().StringSlice("add-host", []string{}, "add a custom host-to-IP mapping (`host:ip`) (default [])")
	farm_buildCmd.Flags().Bool("all-platforms", false, "attempt to build for all base image platforms")
	farm_buildCmd.Flags().StringSlice("annotation", []string{}, "set metadata for an image (default [])")
	farm_buildCmd.Flags().String("arch", "", "set the ARCH of the image to the provided value instead of the architecture of the host")
	farm_buildCmd.Flags().String("authfile", "", "path of the authentication file.")
	farm_buildCmd.Flags().String("blob-cache", "", "assume image blobs in the specified directory will be available for pushing")
	farm_buildCmd.Flags().StringSlice("build-arg", []string{}, "`argument=value` to supply to the builder")
	farm_buildCmd.Flags().StringSlice("build-arg-file", []string{}, "`argfile.conf` containing lines of argument=value to supply to the builder")
	farm_buildCmd.Flags().StringSlice("build-context", []string{}, "`argument=value` to supply additional build context to the builder")
	farm_buildCmd.Flags().StringSlice("cache-from", []string{}, "remote repository list to utilise as potential cache source.")
	farm_buildCmd.Flags().StringSlice("cache-to", []string{}, "remote repository list to utilise as potential cache destination.")
	farm_buildCmd.Flags().String("cache-ttl", "", "only consider cache images under specified duration.")
	farm_buildCmd.Flags().StringSlice("cap-add", []string{}, "add the specified capability when running (default [])")
	farm_buildCmd.Flags().StringSlice("cap-drop", []string{}, "drop the specified capability when running (default [])")
	farm_buildCmd.Flags().String("cdi-config-dir", "", "`directory` of CDI configuration files")
	farm_buildCmd.Flags().String("cert-dir", "", "use certificates at the specified path to access the registry")
	farm_buildCmd.Flags().String("cgroup-parent", "", "optional parent cgroup for the container")
	farm_buildCmd.Flags().String("cgroupns", "", "'private', or 'host'")
	farm_buildCmd.Flags().Bool("cleanup", false, "Remove built images from farm nodes on success")
	farm_buildCmd.Flags().String("cni-config-dir", "", "`directory` of CNI configuration files")
	farm_buildCmd.Flags().String("cni-plugin-path", "", "`path` of CNI network plugins")
	farm_buildCmd.Flags().Bool("compat-volumes", false, "preserve the contents of VOLUMEs during RUN instructions")
	farm_buildCmd.Flags().Bool("compress", false, "this is a legacy option, which has no effect on the image")
	farm_buildCmd.Flags().StringSlice("cpp-flag", []string{}, "set additional flag to pass to C preprocessor (cpp)")
	farm_buildCmd.Flags().String("cpu-period", "", "limit the CPU CFS (Completely Fair Scheduler) period")
	farm_buildCmd.Flags().String("cpu-quota", "", "limit the CPU CFS (Completely Fair Scheduler) quota")
	farm_buildCmd.Flags().StringP("cpu-shares", "c", "", "CPU shares (relative weight)")
	farm_buildCmd.Flags().String("cpuset-cpus", "", "CPUs in which to allow execution (0-3, 0,1)")
	farm_buildCmd.Flags().String("cpuset-mems", "", "memory nodes (MEMs) in which to allow execution (0-3, 0,1). Only effective on NUMA systems.")
	farm_buildCmd.Flags().String("creds", "", "use `[username[:password]]` for accessing the registry")
	farm_buildCmd.Flags().String("cw", "", "confidential workload `options`")
	farm_buildCmd.Flags().StringSlice("decryption-key", []string{}, "key needed to decrypt the image")
	farm_buildCmd.Flags().StringSlice("device", []string{}, "additional devices to provide")
	farm_buildCmd.Flags().BoolP("disable-compression", "D", false, "don't compress layers by default")
	farm_buildCmd.Flags().Bool("disable-content-trust", false, "this is a Docker specific option and is a NOOP")
	farm_buildCmd.Flags().StringSlice("dns", []string{}, "set custom DNS servers or disable it completely by setting it to 'none', which prevents the automatic creation of `/etc/resolv.conf`.")
	farm_buildCmd.Flags().StringSlice("dns-option", []string{}, "set custom DNS options")
	farm_buildCmd.Flags().StringSlice("dns-search", []string{}, "set custom DNS search domains")
	farm_buildCmd.Flags().StringSlice("env", []string{}, "set environment variable for the image")
	farm_buildCmd.Flags().String("farm", "", "Farm to use for builds")
	farm_buildCmd.Flags().StringSliceP("file", "f", []string{}, "`pathname or URL` of a Dockerfile")
	farm_buildCmd.Flags().Bool("force-rm", false, "always remove intermediate containers after a build, even if the build is unsuccessful.")
	farm_buildCmd.Flags().String("format", "", "`format` of the built image's manifest and metadata. Use BUILDAH_FORMAT environment variable to override.")
	farm_buildCmd.Flags().String("from", "", "image name used to replace the value in the first FROM instruction in the Containerfile")
	farm_buildCmd.Flags().StringSlice("group-add", []string{}, "add additional groups to the primary container process. 'keep-groups' allows container processes to use supplementary groups.")
	farm_buildCmd.Flags().StringSlice("hooks-dir", []string{}, "set the OCI hooks directory path (may be set multiple times)")
	farm_buildCmd.Flags().Bool("http-proxy", false, "pass through HTTP Proxy environment variables")
	farm_buildCmd.Flags().Bool("identity-label", false, "add default identity label")
	farm_buildCmd.Flags().String("ignorefile", "", "path to an alternate .dockerignore file")
	farm_buildCmd.Flags().String("iidfile", "", "`file` to write the image ID to")
	farm_buildCmd.Flags().String("ipc", "", "'private', `path` of IPC namespace to join, or 'host'")
	farm_buildCmd.Flags().String("isolation", "", "`type` of process isolation to use. Use BUILDAH_ISOLATION environment variable to override.")
	farm_buildCmd.Flags().String("jobs", "", "how many stages to run in parallel")
	farm_buildCmd.Flags().StringSlice("label", []string{}, "set metadata for an image (default [])")
	farm_buildCmd.Flags().StringSlice("layer-label", []string{}, "set metadata for an intermediate image (default [])")
	farm_buildCmd.Flags().Bool("layers", false, "use intermediate layers during build. Use BUILDAH_LAYERS environment variable to override.")
	farm_buildCmd.Flags().Bool("load", false, "buildx --load")
	farm_buildCmd.Flags().BoolP("local", "l", false, "Build image on local machine as well as on farm nodes")
	farm_buildCmd.Flags().Bool("log-rusage", false, "log resource usage at each build step")
	farm_buildCmd.Flags().String("logfile", "", "log to `file` instead of stdout/stderr")
	farm_buildCmd.Flags().String("loglevel", "", "NO LONGER USED, flag ignored, and hidden")
	farm_buildCmd.Flags().Bool("logsplit", false, "split logfile to different files for each platform")
	farm_buildCmd.Flags().String("manifest", "", "add the image to the specified manifest list. Creates manifest list if it does not exist")
	farm_buildCmd.Flags().StringP("memory", "m", "", "memory limit (format: <number>[<unit>], where unit = b, k, m or g)")
	farm_buildCmd.Flags().String("memory-swap", "", "swap limit equal to memory plus swap: '-1' to enable unlimited swap")
	farm_buildCmd.Flags().String("network", "", "'private', 'none', 'ns:path' of network namespace to join, or 'host'")
	farm_buildCmd.Flags().Bool("no-cache", false, "do not use existing cached images for the container build. Build from the start with a new set of cached layers.")
	farm_buildCmd.Flags().Bool("no-hostname", false, "do not create new /etc/hostname file for RUN instructions, use the one from the base image.")
	farm_buildCmd.Flags().Bool("no-hosts", false, "do not create new /etc/hosts file for RUN instructions, use the one from the base image.")
	farm_buildCmd.Flags().Bool("omit-history", false, "omit build history information from built image")
	farm_buildCmd.Flags().String("os", "", "set the OS to the provided value instead of the current operating system of the host")
	farm_buildCmd.Flags().StringSlice("os-feature", []string{}, "set required OS `feature` for the target image in addition to values from the base image")
	farm_buildCmd.Flags().String("os-version", "", "set required OS `version` for the target image instead of the value from the base image")
	farm_buildCmd.Flags().StringP("output", "o", "", "output destination (format: type=local,dest=path)")
	farm_buildCmd.Flags().String("pid", "", "private, `path` of PID namespace to join, or 'host'")
	farm_buildCmd.Flags().StringSlice("platform", []string{}, "set the `OS/ARCH[/VARIANT]` of the image to the provided value instead of the current operating system and architecture of the host (for example \"linux/arm\")")
	farm_buildCmd.PersistentFlags().StringSlice("platforms", []string{}, "Build only on farm nodes that match the given platforms")
	farm_buildCmd.Flags().String("progress", "", "buildx --progress")
	farm_buildCmd.Flags().String("pull", "", "Pull image policy (\"always\"|\"missing\"|\"never\"|\"newer\")")
	farm_buildCmd.Flags().Bool("pull-always", false, "pull the image even if the named image is present in store")
	farm_buildCmd.Flags().Bool("pull-never", false, "do not pull the image, use the image present in store if available")
	farm_buildCmd.Flags().BoolP("quiet", "q", false, "refrain from announcing build instructions and image read/write progress")
	farm_buildCmd.Flags().String("retry", "", "number of times to retry in case of failure when performing push/pull")
	farm_buildCmd.Flags().String("retry-delay", "", "delay between retries in case of push/pull failures")
	farm_buildCmd.Flags().Bool("rm", false, "remove intermediate containers after a successful build")
	farm_buildCmd.Flags().StringSlice("runtime-flag", []string{}, "add global flags for the container runtime")
	farm_buildCmd.Flags().String("rusage-logfile", "", "destination file to which rusage should be logged to instead of stdout (= the default).")
	farm_buildCmd.Flags().String("sbom", "", "scan working container using `preset` configuration")
	farm_buildCmd.Flags().String("sbom-image-output", "", "add scan results to image as `path`")
	farm_buildCmd.Flags().String("sbom-image-purl-output", "", "add scan results to image as `path`")
	farm_buildCmd.Flags().String("sbom-merge-strategy", "", "merge scan results using `strategy`")
	farm_buildCmd.Flags().String("sbom-output", "", "save scan results to `file`")
	farm_buildCmd.Flags().String("sbom-purl-output", "", "save scan results to `file``")
	farm_buildCmd.Flags().StringSlice("sbom-scanner-command", []string{}, "scan working container using `command` in scanner image")
	farm_buildCmd.Flags().String("sbom-scanner-image", "", "scan working container using scanner command from `image`")
	farm_buildCmd.Flags().StringSlice("secret", []string{}, "secret file to expose to the build")
	farm_buildCmd.Flags().StringSlice("security-opt", []string{}, "security options (default [])")
	farm_buildCmd.Flags().String("shm-size", "", "size of '/dev/shm'. The format is `<number><unit>`.")
	farm_buildCmd.Flags().String("sign-by", "", "sign the image using a GPG key with the specified `FINGERPRINT`")
	farm_buildCmd.Flags().String("signature-policy", "", "`pathname` of signature policy file (not usually used)")
	farm_buildCmd.Flags().Bool("skip-unused-stages", false, "skips stages in multi-stage builds which do not affect the final target")
	farm_buildCmd.Flags().Bool("squash", false, "squash all image layers into a single layer")
	farm_buildCmd.Flags().Bool("squash-all", false, "Squash all layers into a single layer")
	farm_buildCmd.Flags().StringSlice("ssh", []string{}, "SSH agent socket or keys to expose to the build. (format: default|<id>[=<socket>|<key>[,<key>]])")
	farm_buildCmd.Flags().Bool("stdin", false, "pass stdin into containers")
	farm_buildCmd.Flags().StringSliceP("tag", "t", []string{}, "tagged `name` to apply to the built image")
	farm_buildCmd.Flags().String("target", "", "set the target build stage to build")
	farm_buildCmd.Flags().String("timestamp", "", "set created timestamp to the specified epoch seconds to allow for deterministic builds, defaults to current time")
	farm_buildCmd.Flags().Bool("tls-verify", false, "require HTTPS and verify certificates when accessing the registry")
	farm_buildCmd.Flags().StringSlice("ulimit", []string{}, "ulimit options")
	farm_buildCmd.Flags().StringSlice("unsetenv", []string{}, "unset environment variable from final image")
	farm_buildCmd.Flags().StringSlice("unsetlabel", []string{}, "unset label when inheriting labels from base image")
	farm_buildCmd.Flags().String("userns", "", "'container', `path` of user namespace to join, or 'host'")
	farm_buildCmd.Flags().StringSlice("userns-gid-map", []string{}, "`containerGID:hostGID:length` GID mapping to use in user namespace")
	farm_buildCmd.Flags().String("userns-gid-map-group", "", "`name` of entries from /etc/subgid to use to set user namespace GID mapping")
	farm_buildCmd.Flags().StringSlice("userns-uid-map", []string{}, "`containerUID:hostUID:length` UID mapping to use in user namespace")
	farm_buildCmd.Flags().String("userns-uid-map-user", "", "`name` of entries from /etc/subuid to use to set user namespace UID mapping")
	farm_buildCmd.Flags().String("uts", "", "private, :`path` of UTS namespace to join, or 'host'")
	farm_buildCmd.Flags().String("variant", "", "override the `variant` of the specified image")
	farm_buildCmd.Flags().StringSliceP("volume", "v", []string{}, "bind mount a volume into the container")
	farm_buildCmd.Flag("all-platforms").Hidden = true
	farm_buildCmd.Flag("arch").Hidden = true
	farm_buildCmd.Flag("blob-cache").Hidden = true
	farm_buildCmd.Flag("cdi-config-dir").Hidden = true
	farm_buildCmd.Flag("cni-config-dir").Hidden = true
	farm_buildCmd.Flag("cni-plugin-path").Hidden = true
	farm_buildCmd.Flag("compress").Hidden = true
	farm_buildCmd.Flag("cw").Hidden = true
	farm_buildCmd.Flag("disable-content-trust").Hidden = true
	farm_buildCmd.Flag("load").Hidden = true
	farm_buildCmd.Flag("log-rusage").Hidden = true
	farm_buildCmd.Flag("loglevel").Hidden = true
	farm_buildCmd.Flag("logsplit").Hidden = true
	farm_buildCmd.Flag("manifest").Hidden = true
	farm_buildCmd.Flag("os").Hidden = true
	farm_buildCmd.Flag("output").Hidden = true
	farm_buildCmd.Flag("platform").Hidden = true
	farm_buildCmd.Flag("progress").Hidden = true
	farm_buildCmd.Flag("pull").NoOptDefVal = " "
	farm_buildCmd.Flag("pull-always").Hidden = true
	farm_buildCmd.Flag("pull-never").Hidden = true
	farm_buildCmd.Flag("rusage-logfile").Hidden = true
	farm_buildCmd.Flag("sign-by").Hidden = true
	farm_buildCmd.Flag("signature-policy").Hidden = true
	farm_buildCmd.Flag("stdin").Hidden = true
	farm_buildCmd.Flag("variant").Hidden = true
	farmCmd.AddCommand(farm_buildCmd)
}
