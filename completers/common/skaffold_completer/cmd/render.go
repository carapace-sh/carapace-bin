package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var renderCmd = &cobra.Command{
	Use:   "render",
	Short: "Generate rendered Kubernetes manifests",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(renderCmd).Standalone()

	renderCmd.Flags().Bool("assume-yes", false, "If true, skaffold will skip yes/no confirmation from the user and default to yes")
	renderCmd.Flags().StringP("build-artifacts", "a", "", "File containing pre-built images to use instead of rebuilding artifacts. A sample file looks like the following:")
	renderCmd.Flags().Bool("cache-artifacts", true, "Set to false to disable default caching of artifacts")
	renderCmd.Flags().StringP("default-repo", "d", "", "Default repository value (overrides global config)")
	renderCmd.Flags().String("digest-source", "", "Set to 'remote' to skip builds and resolve the digest of images by tag from the remote registry. Set to 'local' to build images locally and use digests from built images. Set to 'tag' to use tags directly from the build. Set to 'none' to use tags directly from the Kubernetes manifests. If unspecified, defaults to 'remote' for remote clusters, and 'tag' for local clusters like kind or minikube.")
	renderCmd.Flags().Bool("enable-gke-arm-node-toleration", true, "Setting this flag provides the appropriate toleration for Pods to be scheduled on GKE ARM nodes that are tainted to disallow workloads by default.")
	renderCmd.Flags().Bool("enable-platform-node-affinity", false, "If true, when deploying to a mixed node cluster, skaffold will add platform (os/arch) node affinity definition to rendered manifests based on the image platforms")
	renderCmd.Flags().Bool("enable-rpc", false, "Enable gRPC or HTTP APIs for exposing Skaffold events")
	renderCmd.Flags().String("event-log-file", "", "Save Skaffold events to the provided file after skaffold has finished executing, requires --rpc-port or --rpc-http-port")
	renderCmd.Flags().StringP("filename", "f", "skaffold.yaml", "Path or URL to the Skaffold config file")
	renderCmd.Flags().String("hydration-dir", ".kpt-pipeline", "The directory to where the (kpt) hydration takes place. Default to a hidden directory .kpt-pipeline.")
	renderCmd.Flags().StringP("images", "i", "", "A list of pre-built images to deploy, either tagged images or NAME=TAG pairs")
	renderCmd.Flags().StringSliceP("label", "l", nil, "Add custom labels to deployed objects. Set multiple times for multiple labels")
	renderCmd.Flags().String("last-log-file", "", "Save Skaffold output to the provided file after skaffold has finished executing, requires --rpc-port or --rpc-http-port (defaults to $HOME/.skaffold/repos)")
	renderCmd.Flags().Bool("loud", false, "Show the build logs and output")
	renderCmd.Flags().StringSliceP("module", "m", nil, "Filter Skaffold configs to only the provided named modules")
	renderCmd.Flags().StringP("namespace", "n", "", "Runs deployments in the specified namespace. When used with 'render' command, renders manifests contain the namespace")
	renderCmd.Flags().Bool("offline", false, "Do not connect to Kubernetes API server for manifest creation and validation. This is helpful when no Kubernetes cluster is available (e.g. GitOps model). No metadata.namespace attribute is injected in this case - the manifest content does not get changed.")
	renderCmd.Flags().StringP("output", "o", "", "File to write rendered manifests to")
	renderCmd.Flags().StringSlice("platform", nil, "The platform to target for the build artifacts")
	renderCmd.Flags().StringSliceP("profile", "p", nil, "Activate profiles by name (prefixed with `-` to disable a profile)")
	renderCmd.Flags().Bool("profile-auto-activation", true, "Set to false to disable profile auto activation")
	renderCmd.Flags().Bool("propagate-profiles", true, "Setting '--propagate-profiles=false' disables propagating profiles set by the '--profile' flag across config dependencies. This mean that only profiles defined directly in the target 'skaffold.yaml' file are activated.")
	renderCmd.Flags().String("remote-cache-dir", "", "Specify the location of the remote cache (default $HOME/.skaffold/remote-cache)")
	renderCmd.Flags().String("resource-selector-rules-file", "", "Path to JSON file specifying the deny list of yaml objects for skaffold to NOT transform with 'image' and 'label' field replacements.  NOTE: this list is additive to skaffold's default denylist and denylist has priority over allowlist")
	renderCmd.Flags().StringSlice("set", nil, "overrides templated manifest fields by provided key-value pairs")
	renderCmd.Flags().String("set-value-file", "", "overrides templated manifest fields by a file containing key-value pairs in .env file format")
	renderCmd.Flags().String("sync-remote-cache", "always", "Controls how Skaffold manages the remote config cache (see `remote-cache-dir`). One of `always` (default), `missing`, or `never`. `always` syncs remote repositories to latest on access. `missing` only clones remote repositories if they do not exist locally. `never` means the user takes responsibility for updating remote repositories.")
	renderCmd.Flags().StringP("tag", "t", "", "The optional custom tag to use for images which overrides the current Tagger configuration")
	renderCmd.Flags().StringP("user", "u", "", "")
	renderCmd.Flags().Bool("wait-for-connection", false, "Blocks ending execution of skaffold until the /v2/events gRPC/HTTP endpoint is hit")
	renderCmd.Flag("enable-gke-arm-node-toleration").Hidden = true
	renderCmd.Flag("enable-rpc").Hidden = true
	renderCmd.Flag("event-log-file").Hidden = true
	renderCmd.Flag("last-log-file").Hidden = true
	renderCmd.Flag("user").Hidden = true
	applyCommonFlagCompletion(renderCmd)
	rootCmd.AddCommand(renderCmd)
}
