package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devCmd = &cobra.Command{
	Use:   "dev",
	Short: "Run a pipeline in development mode",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devCmd).Standalone()

	devCmd.Flags().Bool("assume-yes", false, "If true, skaffold will skip yes/no confirmation from the user and default to yes")
	devCmd.Flags().Bool("auto", false, "Run with an auto-generated skaffold configuration. This will create a temporary `skaffold.yaml` file and kubernetes manifests necessary to run the application")
	devCmd.Flags().Bool("auto-build", true, "When set to false, builds wait for API request instead of running automatically")
	devCmd.Flags().Bool("auto-create-config", true, "If true, skaffold will try to create a config for the user's run if it doesn't find one")
	devCmd.Flags().Bool("auto-deploy", true, "When set to false, deploys wait for API request instead of running automatically")
	devCmd.Flags().Bool("auto-sync", true, "When set to false, syncs wait for API request instead of running automatically")
	devCmd.Flags().String("build-concurrency", "-1", "Number of concurrently running builds. Set to 0 to run all builds in parallel. Doesn't violate build order among dependencies.")
	devCmd.Flags().Bool("cache-artifacts", true, "Set to false to disable default caching of artifacts")
	devCmd.Flags().String("cache-file", "", "Specify the location of the cache file (default $HOME/.skaffold/cache)")
	devCmd.Flags().Bool("check-cluster-node-platforms", true, "When set to true, images are built for the target platforms matching the active kubernetes cluster node platforms. Enabled by default for `dev`, `debug` and `run`")
	devCmd.Flags().Bool("cleanup", true, "Delete deployments after dev or debug mode is interrupted")
	devCmd.Flags().String("cloud-run-location", "", "The GCP Region to deploy Cloud Run services to")
	devCmd.Flags().String("cloud-run-project", "", "The GCP Project ID or Project Number to deploy for Cloud Run")
	devCmd.Flags().StringP("config", "c", "", "File for global configurations (defaults to $HOME/.skaffold/config)")
	devCmd.Flags().StringP("default-repo", "d", "", "Default repository value (overrides global config)")
	devCmd.Flags().Bool("detect-minikube", true, "Use heuristics to detect a minikube cluster")
	devCmd.Flags().String("digest-source", "", "Set to 'remote' to skip builds and resolve the digest of images by tag from the remote registry. Set to 'local' to build images locally and use digests from built images. Set to 'tag' to use tags directly from the build. Set to 'none' to use tags directly from the Kubernetes manifests. If unspecified, defaults to 'remote' for remote clusters, and 'tag' for local clusters like kind or minikube.")
	devCmd.Flags().Bool("disable-multi-platform-build", true, "When set to true, forces only single platform image builds even when multiple target platforms are specified. Enabled by default for `dev` and `debug` modes, to keep dev-loop fast")
	devCmd.Flags().Bool("enable-gke-arm-node-toleration", true, "Setting this flag provides the appropriate toleration for Pods to be scheduled on GKE ARM nodes that are tainted to disallow workloads by default.")
	devCmd.Flags().Bool("enable-platform-node-affinity", true, "If true, when deploying to a mixed node cluster, skaffold will add platform (os/arch) node affinity definition to rendered manifests based on the image platforms")
	devCmd.Flags().Bool("enable-rpc", false, "Enable gRPC or HTTP APIs for exposing Skaffold events")
	devCmd.Flags().String("event-log-file", "", "Save Skaffold events to the provided file after skaffold has finished executing, requires --rpc-port or --rpc-http-port")
	devCmd.Flags().Bool("fast-fail-status-check", true, "Configures `status-check` to fail immediately if any error occurs.  Otherwise `status-check` will attempt to check all resources once and only then report errors and possibly exit.  Defaults to 'true'")
	devCmd.Flags().StringP("filename", "f", "skaffold.yaml", "Path or URL to the Skaffold config file")
	devCmd.Flags().Bool("force", false, "Recreate Kubernetes resources if necessary for deployment, warning: might cause downtime!")
	devCmd.Flags().String("hydration-dir", ".kpt-pipeline", "The directory to where the (kpt) hydration takes place. Default to a hidden directory .kpt-pipeline.")
	devCmd.Flags().StringSlice("insecure-registry", nil, "Target registries for built images which are not secure")
	devCmd.Flags().Bool("iterative-status-check", true, "Run `status-check` iteratively after each deploy step, instead of all-together at the end of all deploys (default).")
	devCmd.Flags().Bool("keep-running-on-failure", false, "If true, the session will be suspended instead of ending if any errors occur, the user can fix the errors during the session suspension, the session can be restored and continued by pressing any key. ")
	devCmd.Flags().String("kube-context", "", "Deploy to this Kubernetes context")
	devCmd.Flags().String("kubeconfig", "", "Path to the kubeconfig file to use for CLI requests.")
	devCmd.Flags().StringSliceP("label", "l", nil, "Add custom labels to deployed objects. Set multiple times for multiple labels")
	devCmd.Flags().String("last-log-file", "", "Save Skaffold output to the provided file after skaffold has finished executing, requires --rpc-port or --rpc-http-port (defaults to $HOME/.skaffold/repos)")
	devCmd.Flags().String("minikube-profile", "", "forces skaffold use the given minikube-profile and forces building against the docker daemon inside that minikube profile")
	devCmd.Flags().StringSliceP("module", "m", nil, "Filter Skaffold configs to only the provided named modules")
	devCmd.Flags().StringSlice("mute-logs", nil, "mute logs for specified stages in pipeline (build, deploy, status-check, none, all)")
	devCmd.Flags().StringP("namespace", "n", "", "Runs deployments in the specified namespace. When used with 'render' command, renders manifests contain the namespace")
	devCmd.Flags().Bool("no-prune", false, "Skip removing images and containers built by Skaffold during cleanup after dev or debug mode")
	devCmd.Flags().Bool("no-prune-children", false, "Skip removing layers reused by Skaffold")
	devCmd.Flags().StringSlice("platform", nil, "The platform to target for the build artifacts")
	devCmd.Flags().StringSlice("port-forward", []string{"user"}, "Port-forward exposes service ports and container ports within pods and other resources (off, user, services, debug, pods)")
	devCmd.Flags().StringSliceP("profile", "p", nil, "Activate profiles by name (prefixed with `-` to disable a profile)")
	devCmd.Flags().Bool("profile-auto-activation", true, "Set to false to disable profile auto activation")
	devCmd.Flags().Bool("propagate-profiles", true, "Setting '--propagate-profiles=false' disables propagating profiles set by the '--profile' flag across config dependencies. This mean that only profiles defined directly in the target 'skaffold.yaml' file are activated.")
	devCmd.Flags().String("remote-cache-dir", "", "Specify the location of the remote cache (default $HOME/.skaffold/remote-cache)")
	devCmd.Flags().Bool("render-only", false, "Print rendered Kubernetes manifests instead of deploying them")
	devCmd.Flags().String("resource-selector-rules-file", "", "Path to JSON file specifying the deny list of yaml objects for skaffold to NOT transform with 'image' and 'label' field replacements.  NOTE: this list is additive to skaffold's default denylist and denylist has priority over allowlist")
	devCmd.Flags().String("rpc-http-port", "", "tcp port to expose the Skaffold API over HTTP REST")
	devCmd.Flags().String("rpc-port", "", "tcp port to expose the Skaffold API over gRPC")
	devCmd.Flags().Bool("skip-tests", false, "Whether to skip the tests after building")
	devCmd.Flags().Bool("status-check", false, "Wait for deployed resources to stabilize")
	devCmd.Flags().String("status-check-selectors", "", "File containing resource selectors for kubernetes resources status check. A sample file looks like the following:")
	devCmd.Flags().String("sync-remote-cache", "always", "Controls how Skaffold manages the remote config cache (see `remote-cache-dir`). One of `always` (default), `missing`, or `never`. `always` syncs remote repositories to latest on access. `missing` only clones remote repositories if they do not exist locally. `never` means the user takes responsibility for updating remote repositories.")
	devCmd.Flags().StringP("tag", "t", "", "The optional custom tag to use for images which overrides the current Tagger configuration")
	devCmd.Flags().Bool("tail", true, "Stream logs from deployed objects")
	devCmd.Flags().Bool("tolerate-failures-until-deadline", false, "Configures `status-check` to tolerate failures until Skaffold's statusCheckDeadline duration or the deployments progressDeadlineSeconds  Otherwise deployment failures skaffold encounters will immediately fail the deployment.  Defaults to 'false'")
	devCmd.Flags().Bool("toot", false, "Emit a terminal beep after the deploy is complete")
	devCmd.Flags().String("trigger", "notify", "How is change detection triggered? (polling, notify, or manual)")
	devCmd.Flags().StringP("user", "u", "", "")
	devCmd.Flags().Bool("wait-for-connection", false, "Blocks ending execution of skaffold until the /v2/events gRPC/HTTP endpoint is hit")
	devCmd.Flags().Bool("wait-for-deletions", true, "Wait for pending deletions to complete before a deployment")
	devCmd.Flags().String("wait-for-deletions-delay", "2s", "Delay between two checks for pending deletions")
	devCmd.Flags().String("wait-for-deletions-max", "1m0s", "Max duration to wait for pending deletions")
	devCmd.Flags().StringSliceP("watch-image", "w", nil, "Choose which artifacts to watch. Artifacts with image names that contain the expression will be watched only. Default is to watch sources for all artifacts")
	devCmd.Flags().StringP("watch-poll-interval", "i", "1000", "Interval (in ms) between two checks for file changes")
	devCmd.Flag("enable-gke-arm-node-toleration").Hidden = true
	devCmd.Flag("enable-rpc").Hidden = true
	devCmd.Flag("event-log-file").Hidden = true
	devCmd.Flag("fast-fail-status-check").Hidden = true
	devCmd.Flag("last-log-file").Hidden = true
	devCmd.Flag("minikube-profile").Hidden = true
	devCmd.Flag("port-forward").NoOptDefVal = " "
	devCmd.Flag("render-only").Hidden = true
	devCmd.Flag("user").Hidden = true
	applyCommonFlagCompletion(devCmd)
	rootCmd.AddCommand(devCmd)
}
