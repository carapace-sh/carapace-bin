package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Apply hydrated manifests to a cluster",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(applyCmd).Standalone()

	applyCmd.Flags().Bool("assume-yes", false, "If true, skaffold will skip yes/no confirmation from the user and default to yes")
	applyCmd.Flags().String("cloud-run-location", "", "The GCP Region to deploy Cloud Run services to")
	applyCmd.Flags().String("cloud-run-project", "", "The GCP Project ID or Project Number to deploy for Cloud Run")
	applyCmd.Flags().StringP("config", "c", "", "File for global configurations (defaults to $HOME/.skaffold/config)")
	applyCmd.Flags().Bool("enable-rpc", false, "Enable gRPC or HTTP APIs for exposing Skaffold events")
	applyCmd.Flags().String("event-log-file", "", "Save Skaffold events to the provided file after skaffold has finished executing, requires --rpc-port or --rpc-http-port")
	applyCmd.Flags().Bool("fast-fail-status-check", true, "Configures `status-check` to fail immediately if any error occurs.  Otherwise `status-check` will attempt to check all resources once and only then report errors and possibly exit.  Defaults to 'true'")
	applyCmd.Flags().StringP("filename", "f", "skaffold.yaml", "Path or URL to the Skaffold config file")
	applyCmd.Flags().Bool("force", false, "Recreate Kubernetes resources if necessary for deployment, warning: might cause downtime!")
	applyCmd.Flags().String("inventory-id", "", "The inventory name, default to `.kpt-pipeline`")
	applyCmd.Flags().String("inventory-name", "", "Inventory identifier.")
	applyCmd.Flags().String("inventory-namespace", "", "The namespace for the ResourceGroup resource that contains the inventory")
	applyCmd.Flags().Bool("iterative-status-check", true, "Run `status-check` iteratively after each deploy step, instead of all-together at the end of all deploys (default).")
	applyCmd.Flags().String("kube-context", "", "Deploy to this Kubernetes context")
	applyCmd.Flags().String("kubeconfig", "", "Path to the kubeconfig file to use for CLI requests.")
	applyCmd.Flags().StringSliceP("label", "l", nil, "Add custom labels to deployed objects. Set multiple times for multiple labels")
	applyCmd.Flags().String("last-log-file", "", "Save Skaffold output to the provided file after skaffold has finished executing, requires --rpc-port or --rpc-http-port (defaults to $HOME/.skaffold/repos)")
	applyCmd.Flags().StringSliceP("module", "m", nil, "Filter Skaffold configs to only the provided named modules")
	applyCmd.Flags().StringP("namespace", "n", "", "Runs deployments in the specified namespace. When used with 'render' command, renders manifests contain the namespace")
	applyCmd.Flags().StringSliceP("profile", "p", nil, "Activate profiles by name (prefixed with `-` to disable a profile)")
	applyCmd.Flags().String("remote-cache-dir", "", "Specify the location of the remote cache (default $HOME/.skaffold/remote-cache)")
	applyCmd.Flags().String("rpc-http-port", "", "tcp port to expose the Skaffold API over HTTP REST")
	applyCmd.Flags().String("rpc-port", "", "tcp port to expose the Skaffold API over gRPC")
	applyCmd.Flags().Bool("status-check", false, "Wait for deployed resources to stabilize")
	applyCmd.Flags().String("status-check-selectors", "", "File containing resource selectors for kubernetes resources status check. A sample file looks like the following:")
	applyCmd.Flags().String("sync-remote-cache", "always", "Controls how Skaffold manages the remote config cache (see `remote-cache-dir`). One of `always` (default), `missing`, or `never`. `always` syncs remote repositories to latest on access. `missing` only clones remote repositories if they do not exist locally. `never` means the user takes responsibility for updating remote repositories.")
	applyCmd.Flags().Bool("tail", false, "Stream logs from deployed objects")
	applyCmd.Flags().Bool("tolerate-failures-until-deadline", false, "Configures `status-check` to tolerate failures until Skaffold's statusCheckDeadline duration or the deployments progressDeadlineSeconds  Otherwise deployment failures skaffold encounters will immediately fail the deployment.  Defaults to 'false'")
	applyCmd.Flags().StringP("user", "u", "", "")
	applyCmd.Flags().Bool("wait-for-connection", false, "Blocks ending execution of skaffold until the /v2/events gRPC/HTTP endpoint is hit")
	applyCmd.Flag("enable-rpc").Hidden = true
	applyCmd.Flag("event-log-file").Hidden = true
	applyCmd.Flag("fast-fail-status-check").Hidden = true
	applyCmd.Flag("inventory-id").Hidden = true
	applyCmd.Flag("inventory-name").Hidden = true
	applyCmd.Flag("inventory-namespace").Hidden = true
	applyCmd.Flag("last-log-file").Hidden = true
	applyCmd.Flag("user").Hidden = true
	applyCommonFlagCompletion(applyCmd)
	rootCmd.AddCommand(applyCmd)
}
