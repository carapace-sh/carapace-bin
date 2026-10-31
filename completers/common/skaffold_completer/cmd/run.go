package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a pipeline",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runCmd).Standalone()

	runCmd.Flags().Bool("assume-yes", false, "If true, skaffold will skip yes/no confirmation from the user and default to yes")
	runCmd.Flags().Bool("auto", false, "Run with an auto-generated skaffold configuration. This will create a temporary `skaffold.yaml` file and kubernetes manifests necessary to run the application")
	runCmd.Flags().Bool("auto-create-config", true, "If true, skaffold will try to create a config for the user's run if it doesn't find one")
	runCmd.Flags().String("build-concurrency", "-1", "Number of concurrently running builds. Set to 0 to run all builds in parallel. Doesn't violate build order among dependencies.")
	runCmd.Flags().StringSliceP("build-image", "b", nil, "Only build artifacts with image names that contain the given substring. Default is to build sources for all artifacts")
	runCmd.Flags().Bool("cache-artifacts", true, "Set to false to disable default caching of artifacts")
	runCmd.Flags().String("cache-file", "", "Specify the location of the cache file (default $HOME/.skaffold/cache)")
	runCmd.Flags().Bool("check-cluster-node-platforms", true, "When set to true, images are built for the target platforms matching the active kubernetes cluster node platforms. Enabled by default for `dev`, `debug` and `run`")
	runCmd.Flags().Bool("cleanup", true, "Delete deployments after dev or debug mode is interrupted")
	runCmd.Flags().String("cloud-run-location", "", "The GCP Region to deploy Cloud Run services to")
	runCmd.Flags().String("cloud-run-project", "", "The GCP Project ID or Project Number to deploy for Cloud Run")
	runCmd.Flags().StringP("config", "c", "", "File for global configurations (defaults to $HOME/.skaffold/config)")
	runCmd.Flags().StringP("default-repo", "d", "", "Default repository value (overrides global config)")
	runCmd.Flags().Bool("detect-minikube", true, "Use heuristics to detect a minikube cluster")
	runCmd.Flags().String("digest-source", "", "Set to 'remote' to skip builds and resolve the digest of images by tag from the remote registry. Set to 'local' to build images locally and use digests from built images. Set to 'tag' to use tags directly from the build. Set to 'none' to use tags directly from the Kubernetes manifests. If unspecified, defaults to 'remote' for remote clusters, and 'tag' for local clusters like kind or minikube.")
	runCmd.Flags().Bool("disable-multi-platform-build", false, "When set to true, forces only single platform image builds even when multiple target platforms are specified. Enabled by default for `dev` and `debug` modes, to keep dev-loop fast")
	runCmd.Flags().Bool("enable-gke-arm-node-toleration", true, "Setting this flag provides the appropriate toleration for Pods to be scheduled on GKE ARM nodes that are tainted to disallow workloads by default.")
	runCmd.Flags().Bool("enable-platform-node-affinity", true, "If true, when deploying to a mixed node cluster, skaffold will add platform (os/arch) node affinity definition to rendered manifests based on the image platforms")
	runCmd.Flags().Bool("enable-rpc", false, "Enable gRPC or HTTP APIs for exposing Skaffold events")
	runCmd.Flags().String("event-log-file", "", "Save Skaffold events to the provided file after skaffold has finished executing, requires --rpc-port or --rpc-http-port")
	runCmd.Flags().Bool("fast-fail-status-check", true, "Configures `status-check` to fail immediately if any error occurs.  Otherwise `status-check` will attempt to check all resources once and only then report errors and possibly exit.  Defaults to 'true'")
	runCmd.Flags().StringP("filename", "f", "skaffold.yaml", "Path or URL to the Skaffold config file")
	runCmd.Flags().Bool("force", false, "Recreate Kubernetes resources if necessary for deployment, warning: might cause downtime!")
	runCmd.Flags().String("hydration-dir", ".kpt-pipeline", "The directory to where the (kpt) hydration takes place. Default to a hidden directory .kpt-pipeline.")
	runCmd.Flags().StringSlice("insecure-registry", nil, "Target registries for built images which are not secure")
	runCmd.Flags().Bool("iterative-status-check", true, "Run `status-check` iteratively after each deploy step, instead of all-together at the end of all deploys (default).")
	runCmd.Flags().String("kube-context", "", "Deploy to this Kubernetes context")
	runCmd.Flags().String("kubeconfig", "", "Path to the kubeconfig file to use for CLI requests.")
	runCmd.Flags().StringSliceP("label", "l", nil, "Add custom labels to deployed objects. Set multiple times for multiple labels")
	runCmd.Flags().String("last-log-file", "", "Save Skaffold output to the provided file after skaffold has finished executing, requires --rpc-port or --rpc-http-port (defaults to $HOME/.skaffold/repos)")
	runCmd.Flags().String("minikube-profile", "", "forces skaffold use the given minikube-profile and forces building against the docker daemon inside that minikube profile")
	runCmd.Flags().StringSliceP("module", "m", nil, "Filter Skaffold configs to only the provided named modules")
	runCmd.Flags().StringSlice("mute-logs", nil, "mute logs for specified stages in pipeline (build, deploy, status-check, none, all)")
	runCmd.Flags().StringP("namespace", "n", "", "Runs deployments in the specified namespace. When used with 'render' command, renders manifests contain the namespace")
	runCmd.Flags().Bool("no-prune", false, "Skip removing images and containers built by Skaffold during cleanup after dev or debug mode")
	runCmd.Flags().Bool("no-prune-children", false, "Skip removing layers reused by Skaffold")
	runCmd.Flags().StringSlice("platform", nil, "The platform to target for the build artifacts")
	runCmd.Flags().StringSlice("port-forward", []string{"off"}, "Port-forward exposes service ports and container ports within pods and other resources (off, user, services, debug, pods)")
	runCmd.Flags().StringSliceP("profile", "p", nil, "Activate profiles by name (prefixed with `-` to disable a profile)")
	runCmd.Flags().Bool("profile-auto-activation", true, "Set to false to disable profile auto activation")
	runCmd.Flags().Bool("propagate-profiles", true, "Setting '--propagate-profiles=false' disables propagating profiles set by the '--profile' flag across config dependencies. This mean that only profiles defined directly in the target 'skaffold.yaml' file are activated.")
	runCmd.Flags().String("remote-cache-dir", "", "Specify the location of the remote cache (default $HOME/.skaffold/remote-cache)")
	runCmd.Flags().Bool("render-only", false, "Print rendered Kubernetes manifests instead of deploying them")
	runCmd.Flags().String("render-output", "", "Writes '--render-only' output to the specified file")
	runCmd.Flags().String("resource-selector-rules-file", "", "Path to JSON file specifying the deny list of yaml objects for skaffold to NOT transform with 'image' and 'label' field replacements.  NOTE: this list is additive to skaffold's default denylist and denylist has priority over allowlist")
	runCmd.Flags().String("rpc-http-port", "", "tcp port to expose the Skaffold API over HTTP REST")
	runCmd.Flags().String("rpc-port", "", "tcp port to expose the Skaffold API over gRPC")
	runCmd.Flags().Bool("skip-tests", false, "Whether to skip the tests after building")
	runCmd.Flags().Bool("status-check", false, "Wait for deployed resources to stabilize")
	runCmd.Flags().String("status-check-selectors", "", "File containing resource selectors for kubernetes resources status check. A sample file looks like the following:")
	runCmd.Flags().String("sync-remote-cache", "always", "Controls how Skaffold manages the remote config cache (see `remote-cache-dir`). One of `always` (default), `missing`, or `never`. `always` syncs remote repositories to latest on access. `missing` only clones remote repositories if they do not exist locally. `never` means the user takes responsibility for updating remote repositories.")
	runCmd.Flags().StringP("tag", "t", "", "The optional custom tag to use for images which overrides the current Tagger configuration")
	runCmd.Flags().Bool("tail", false, "Stream logs from deployed objects")
	runCmd.Flags().Bool("tolerate-failures-until-deadline", false, "Configures `status-check` to tolerate failures until Skaffold's statusCheckDeadline duration or the deployments progressDeadlineSeconds  Otherwise deployment failures skaffold encounters will immediately fail the deployment.  Defaults to 'false'")
	runCmd.Flags().Bool("toot", false, "Emit a terminal beep after the deploy is complete")
	runCmd.Flags().StringP("user", "u", "", "")
	runCmd.Flags().Bool("wait-for-connection", false, "Blocks ending execution of skaffold until the /v2/events gRPC/HTTP endpoint is hit")
	runCmd.Flags().Bool("wait-for-deletions", true, "Wait for pending deletions to complete before a deployment")
	runCmd.Flags().String("wait-for-deletions-delay", "2s", "Delay between two checks for pending deletions")
	runCmd.Flags().String("wait-for-deletions-max", "1m0s", "Max duration to wait for pending deletions")
	runCmd.Flag("enable-gke-arm-node-toleration").Hidden = true
	runCmd.Flag("enable-rpc").Hidden = true
	runCmd.Flag("event-log-file").Hidden = true
	runCmd.Flag("fast-fail-status-check").Hidden = true
	runCmd.Flag("last-log-file").Hidden = true
	runCmd.Flag("minikube-profile").Hidden = true
	runCmd.Flag("port-forward").NoOptDefVal = " "
	runCmd.Flag("render-only").Hidden = true
	runCmd.Flag("render-output").Hidden = true
	runCmd.Flag("user").Hidden = true
	applyCommonFlagCompletion(runCmd)
	rootCmd.AddCommand(runCmd)
}
