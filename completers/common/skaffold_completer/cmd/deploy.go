package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy pre-built artifacts",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deployCmd).Standalone()

	deployCmd.Flags().Bool("assume-yes", false, "If true, skaffold will skip yes/no confirmation from the user and default to yes")
	deployCmd.Flags().StringP("build-artifacts", "a", "", "File containing pre-built images to use instead of rebuilding artifacts. A sample file looks like the following:")
	deployCmd.Flags().String("build-concurrency", "-1", "Number of concurrently running builds. Set to 0 to run all builds in parallel. Doesn't violate build order among dependencies.")
	deployCmd.Flags().String("cloud-run-location", "", "The GCP Region to deploy Cloud Run services to")
	deployCmd.Flags().String("cloud-run-project", "", "The GCP Project ID or Project Number to deploy for Cloud Run")
	deployCmd.Flags().StringP("config", "c", "", "File for global configurations (defaults to $HOME/.skaffold/config)")
	deployCmd.Flags().StringP("default-repo", "d", "", "Default repository value (overrides global config)")
	deployCmd.Flags().Bool("detect-minikube", true, "Use heuristics to detect a minikube cluster")
	deployCmd.Flags().Bool("enable-gke-arm-node-toleration", true, "Setting this flag provides the appropriate toleration for Pods to be scheduled on GKE ARM nodes that are tainted to disallow workloads by default.")
	deployCmd.Flags().Bool("enable-platform-node-affinity", false, "If true, when deploying to a mixed node cluster, skaffold will add platform (os/arch) node affinity definition to rendered manifests based on the image platforms")
	deployCmd.Flags().Bool("enable-rpc", false, "Enable gRPC or HTTP APIs for exposing Skaffold events")
	deployCmd.Flags().String("event-log-file", "", "Save Skaffold events to the provided file after skaffold has finished executing, requires --rpc-port or --rpc-http-port")
	deployCmd.Flags().Bool("fast-fail-status-check", true, "Configures `status-check` to fail immediately if any error occurs.  Otherwise `status-check` will attempt to check all resources once and only then report errors and possibly exit.  Defaults to 'true'")
	deployCmd.Flags().StringP("filename", "f", "skaffold.yaml", "Path or URL to the Skaffold config file")
	deployCmd.Flags().Bool("force", false, "Recreate Kubernetes resources if necessary for deployment, warning: might cause downtime!")
	deployCmd.Flags().String("hydration-dir", ".kpt-pipeline", "The directory to where the (kpt) hydration takes place. Default to a hidden directory .kpt-pipeline.")
	deployCmd.Flags().StringP("images", "i", "", "A list of pre-built images to deploy, either tagged images or NAME=TAG pairs")
	deployCmd.Flags().Bool("iterative-status-check", true, "Run `status-check` iteratively after each deploy step, instead of all-together at the end of all deploys (default).")
	deployCmd.Flags().String("kube-context", "", "Deploy to this Kubernetes context")
	deployCmd.Flags().String("kubeconfig", "", "Path to the kubeconfig file to use for CLI requests.")
	deployCmd.Flags().StringSliceP("label", "l", nil, "Add custom labels to deployed objects. Set multiple times for multiple labels")
	deployCmd.Flags().String("last-log-file", "", "Save Skaffold output to the provided file after skaffold has finished executing, requires --rpc-port or --rpc-http-port (defaults to $HOME/.skaffold/repos)")
	deployCmd.Flags().Bool("load-images", false, "If true, skaffold will force load the container images into the local cluster.")
	deployCmd.Flags().StringSliceP("module", "m", nil, "Filter Skaffold configs to only the provided named modules")
	deployCmd.Flags().StringSlice("mute-logs", nil, "mute logs for specified stages in pipeline (build, deploy, status-check, none, all)")
	deployCmd.Flags().StringP("namespace", "n", "", "Runs deployments in the specified namespace. When used with 'render' command, renders manifests contain the namespace")
	deployCmd.Flags().StringSlice("port-forward", []string{"off"}, "Port-forward exposes service ports and container ports within pods and other resources (off, user, services, debug, pods)")
	deployCmd.Flags().StringSliceP("profile", "p", nil, "Activate profiles by name (prefixed with `-` to disable a profile)")
	deployCmd.Flags().Bool("profile-auto-activation", true, "Set to false to disable profile auto activation")
	deployCmd.Flags().Bool("propagate-profiles", true, "Setting '--propagate-profiles=false' disables propagating profiles set by the '--profile' flag across config dependencies. This mean that only profiles defined directly in the target 'skaffold.yaml' file are activated.")
	deployCmd.Flags().String("remote-cache-dir", "", "Specify the location of the remote cache (default $HOME/.skaffold/remote-cache)")
	deployCmd.Flags().String("resource-selector-rules-file", "", "Path to JSON file specifying the deny list of yaml objects for skaffold to NOT transform with 'image' and 'label' field replacements.  NOTE: this list is additive to skaffold's default denylist and denylist has priority over allowlist")
	deployCmd.Flags().String("rpc-http-port", "", "tcp port to expose the Skaffold API over HTTP REST")
	deployCmd.Flags().String("rpc-port", "", "tcp port to expose the Skaffold API over gRPC")
	deployCmd.Flags().Bool("status-check", false, "Wait for deployed resources to stabilize")
	deployCmd.Flags().String("status-check-selectors", "", "File containing resource selectors for kubernetes resources status check. A sample file looks like the following:")
	deployCmd.Flags().String("sync-remote-cache", "always", "Controls how Skaffold manages the remote config cache (see `remote-cache-dir`). One of `always` (default), `missing`, or `never`. `always` syncs remote repositories to latest on access. `missing` only clones remote repositories if they do not exist locally. `never` means the user takes responsibility for updating remote repositories.")
	deployCmd.Flags().StringP("tag", "t", "", "The optional custom tag to use for images which overrides the current Tagger configuration")
	deployCmd.Flags().Bool("tail", false, "Stream logs from deployed objects")
	deployCmd.Flags().Bool("tolerate-failures-until-deadline", false, "Configures `status-check` to tolerate failures until Skaffold's statusCheckDeadline duration or the deployments progressDeadlineSeconds  Otherwise deployment failures skaffold encounters will immediately fail the deployment.  Defaults to 'false'")
	deployCmd.Flags().Bool("toot", false, "Emit a terminal beep after the deploy is complete")
	deployCmd.Flags().StringP("user", "u", "", "")
	deployCmd.Flags().Bool("wait-for-connection", false, "Blocks ending execution of skaffold until the /v2/events gRPC/HTTP endpoint is hit")
	deployCmd.Flags().Bool("wait-for-deletions", true, "Wait for pending deletions to complete before a deployment")
	deployCmd.Flags().String("wait-for-deletions-delay", "2s", "Delay between two checks for pending deletions")
	deployCmd.Flags().String("wait-for-deletions-max", "1m0s", "Max duration to wait for pending deletions")
	deployCmd.Flag("enable-gke-arm-node-toleration").Hidden = true
	deployCmd.Flag("enable-rpc").Hidden = true
	deployCmd.Flag("event-log-file").Hidden = true
	deployCmd.Flag("fast-fail-status-check").Hidden = true
	deployCmd.Flag("last-log-file").Hidden = true
	deployCmd.Flag("port-forward").NoOptDefVal = " "
	deployCmd.Flag("user").Hidden = true
	applyCommonFlagCompletion(deployCmd)
	rootCmd.AddCommand(deployCmd)
}
