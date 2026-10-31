package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build the artifacts",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(buildCmd).Standalone()

	buildCmd.Flags().Bool("assume-yes", false, "If true, skaffold will skip yes/no confirmation from the user and default to yes")
	buildCmd.Flags().String("build-concurrency", "-1", "Number of concurrently running builds. Set to 0 to run all builds in parallel. Doesn't violate build order among dependencies.")
	buildCmd.Flags().StringSliceP("build-image", "b", nil, "Only build artifacts with image names that contain the given substring. Default is to build sources for all artifacts")
	buildCmd.Flags().Bool("cache-artifacts", true, "Set to false to disable default caching of artifacts")
	buildCmd.Flags().String("cache-file", "", "Specify the location of the cache file (default $HOME/.skaffold/cache)")
	buildCmd.Flags().Bool("check-cluster-node-platforms", false, "When set to true, images are built for the target platforms matching the active kubernetes cluster node platforms. Enabled by default for `dev`, `debug` and `run`")
	buildCmd.Flags().StringP("config", "c", "", "File for global configurations (defaults to $HOME/.skaffold/config)")
	buildCmd.Flags().StringP("default-repo", "d", "", "Default repository value (overrides global config)")
	buildCmd.Flags().Bool("detect-minikube", true, "Use heuristics to detect a minikube cluster")
	buildCmd.Flags().Bool("disable-multi-platform-build", false, "When set to true, forces only single platform image builds even when multiple target platforms are specified. Enabled by default for `dev` and `debug` modes, to keep dev-loop fast")
	buildCmd.Flags().Bool("dry-run", false, "Don't build images, just compute the tag for each artifact.")
	buildCmd.Flags().Bool("enable-rpc", false, "Enable gRPC or HTTP APIs for exposing Skaffold events")
	buildCmd.Flags().String("event-log-file", "", "Save Skaffold events to the provided file after skaffold has finished executing, requires --rpc-port or --rpc-http-port")
	buildCmd.Flags().String("file-output", "", "Filename to write build images to")
	buildCmd.Flags().StringP("filename", "f", "skaffold.yaml", "Path or URL to the Skaffold config file")
	buildCmd.Flags().StringSlice("insecure-registry", nil, "Target registries for built images which are not secure")
	buildCmd.Flags().String("kube-context", "", "Deploy to this Kubernetes context")
	buildCmd.Flags().String("kubeconfig", "", "Path to the kubeconfig file to use for CLI requests.")
	buildCmd.Flags().String("last-log-file", "", "Save Skaffold output to the provided file after skaffold has finished executing, requires --rpc-port or --rpc-http-port (defaults to $HOME/.skaffold/repos)")
	buildCmd.Flags().String("minikube-profile", "", "forces skaffold use the given minikube-profile and forces building against the docker daemon inside that minikube profile")
	buildCmd.Flags().StringSliceP("module", "m", nil, "Filter Skaffold configs to only the provided named modules")
	buildCmd.Flags().StringSlice("mute-logs", nil, "mute logs for specified stages in pipeline (build, deploy, status-check, none, all)")
	buildCmd.Flags().StringP("namespace", "n", "", "Runs deployments in the specified namespace. When used with 'render' command, renders manifests contain the namespace")
	buildCmd.Flags().StringP("output", "o", "{{json .}}", "Used in conjunction with --quiet flag. Format output with go-template. For full struct documentation, see https://godoc.org/github.com/GoogleContainerTools/skaffold/v2/cmd/skaffold/app/flags#BuildOutput")
	buildCmd.Flags().StringSlice("platform", nil, "The platform to target for the build artifacts")
	buildCmd.Flags().StringSliceP("profile", "p", nil, "Activate profiles by name (prefixed with `-` to disable a profile)")
	buildCmd.Flags().Bool("profile-auto-activation", true, "Set to false to disable profile auto activation")
	buildCmd.Flags().Bool("propagate-profiles", true, "Setting '--propagate-profiles=false' disables propagating profiles set by the '--profile' flag across config dependencies. This mean that only profiles defined directly in the target 'skaffold.yaml' file are activated.")
	buildCmd.Flags().Bool("push", false, "Push the built images to the specified image repository.")
	buildCmd.Flags().BoolP("quiet", "q", false, "Suppress the build output and print image built on success. See --output to format output.")
	buildCmd.Flags().String("remote-cache-dir", "", "Specify the location of the remote cache (default $HOME/.skaffold/remote-cache)")
	buildCmd.Flags().String("rpc-http-port", "", "tcp port to expose the Skaffold API over HTTP REST")
	buildCmd.Flags().String("rpc-port", "", "tcp port to expose the Skaffold API over gRPC")
	buildCmd.Flags().Bool("skip-tests", false, "Whether to skip the tests after building")
	buildCmd.Flags().String("sync-remote-cache", "always", "Controls how Skaffold manages the remote config cache (see `remote-cache-dir`). One of `always` (default), `missing`, or `never`. `always` syncs remote repositories to latest on access. `missing` only clones remote repositories if they do not exist locally. `never` means the user takes responsibility for updating remote repositories.")
	buildCmd.Flags().StringP("tag", "t", "", "The optional custom tag to use for images which overrides the current Tagger configuration")
	buildCmd.Flags().Bool("toot", false, "Emit a terminal beep after the deploy is complete")
	buildCmd.Flags().StringP("user", "u", "", "")
	buildCmd.Flags().Bool("wait-for-connection", false, "Blocks ending execution of skaffold until the /v2/events gRPC/HTTP endpoint is hit")
	buildCmd.Flag("enable-rpc").Hidden = true
	buildCmd.Flag("event-log-file").Hidden = true
	buildCmd.Flag("last-log-file").Hidden = true
	buildCmd.Flag("minikube-profile").Hidden = true
	buildCmd.Flag("user").Hidden = true
	applyCommonFlagCompletion(buildCmd)
	rootCmd.AddCommand(buildCmd)
}
