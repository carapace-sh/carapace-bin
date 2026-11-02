package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete any resources deployed by Skaffold",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deleteCmd).Standalone()

	deleteCmd.Flags().Bool("assume-yes", false, "If true, skaffold will skip yes/no confirmation from the user and default to yes")
	deleteCmd.Flags().String("cloud-run-location", "", "The GCP Region to deploy Cloud Run services to")
	deleteCmd.Flags().String("cloud-run-project", "", "The GCP Project ID or Project Number to deploy for Cloud Run")
	deleteCmd.Flags().StringP("config", "c", "", "File for global configurations (defaults to $HOME/.skaffold/config)")
	deleteCmd.Flags().StringP("default-repo", "d", "", "Default repository value (overrides global config)")
	deleteCmd.Flags().Bool("detect-minikube", true, "Use heuristics to detect a minikube cluster")
	deleteCmd.Flags().Bool("dry-run", false, "Don't delete resources, just print them.")
	deleteCmd.Flags().StringP("filename", "f", "skaffold.yaml", "Path or URL to the Skaffold config file")
	deleteCmd.Flags().String("kube-context", "", "Deploy to this Kubernetes context")
	deleteCmd.Flags().String("kubeconfig", "", "Path to the kubeconfig file to use for CLI requests.")
	deleteCmd.Flags().StringSliceP("module", "m", nil, "Filter Skaffold configs to only the provided named modules")
	deleteCmd.Flags().StringP("namespace", "n", "", "Runs deployments in the specified namespace. When used with 'render' command, renders manifests contain the namespace")
	deleteCmd.Flags().StringSliceP("profile", "p", nil, "Activate profiles by name (prefixed with `-` to disable a profile)")
	deleteCmd.Flags().Bool("profile-auto-activation", true, "Set to false to disable profile auto activation")
	deleteCmd.Flags().Bool("propagate-profiles", true, "Setting '--propagate-profiles=false' disables propagating profiles set by the '--profile' flag across config dependencies. This mean that only profiles defined directly in the target 'skaffold.yaml' file are activated.")
	deleteCmd.Flags().String("remote-cache-dir", "", "Specify the location of the remote cache (default $HOME/.skaffold/remote-cache)")
	deleteCmd.Flags().StringSlice("set", nil, "overrides templated manifest fields by provided key-value pairs")
	deleteCmd.Flags().String("set-value-file", "", "overrides templated manifest fields by a file containing key-value pairs in .env file format")
	deleteCmd.Flags().String("sync-remote-cache", "always", "Controls how Skaffold manages the remote config cache (see `remote-cache-dir`). One of `always` (default), `missing`, or `never`. `always` syncs remote repositories to latest on access. `missing` only clones remote repositories if they do not exist locally. `never` means the user takes responsibility for updating remote repositories.")
	deleteCmd.Flags().StringP("user", "u", "", "")
	deleteCmd.Flag("user").Hidden = true
	applyCommonFlagCompletion(deleteCmd)
	rootCmd.AddCommand(deleteCmd)
}
