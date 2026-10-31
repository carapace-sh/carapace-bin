package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var filterCmd = &cobra.Command{
	Use:    "filter",
	Short:  "Filter and transform a set of Kubernetes manifests from stdin",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(filterCmd).Standalone()

	filterCmd.Flags().Bool("assume-yes", false, "If true, skaffold will skip yes/no confirmation from the user and default to yes")
	filterCmd.Flags().StringP("build-artifacts", "a", "", "File containing build result from a previous 'skaffold build --file-output'")
	filterCmd.Flags().Bool("debugging", false, "Apply debug transforms similar to \"skaffold debug\"")
	filterCmd.Flags().StringP("filename", "f", "skaffold.yaml", "Path or URL to the Skaffold config file")
	filterCmd.Flags().String("kube-context", "", "Deploy to this Kubernetes context")
	filterCmd.Flags().String("kubeconfig", "", "Path to the kubeconfig file to use for CLI requests.")
	filterCmd.Flags().StringSliceP("label", "l", nil, "Add custom labels to deployed objects. Set multiple times for multiple labels")
	filterCmd.Flags().StringSliceP("module", "m", nil, "Filter Skaffold configs to only the provided named modules")
	filterCmd.Flags().String("post-renderer", "", "Any executable that accepts rendered Kubernetes manifests on STDIN and returns valid Kubernetes manifests on STDOUT")
	filterCmd.Flags().StringSlice("protocols", nil, "Priority sorted order of debugger protocols to support.")
	filterCmd.Flags().String("remote-cache-dir", "", "Specify the location of the remote cache (default $HOME/.skaffold/remote-cache)")
	filterCmd.Flags().StringSlice("set", nil, "overrides templated manifest fields by provided key-value pairs")
	filterCmd.Flags().String("sync-remote-cache", "always", "Controls how Skaffold manages the remote config cache (see `remote-cache-dir`). One of `always` (default), `missing`, or `never`. `always` syncs remote repositories to latest on access. `missing` only clones remote repositories if they do not exist locally. `never` means the user takes responsibility for updating remote repositories.")
	filterCmd.Flags().StringP("user", "u", "", "")
	filterCmd.Flag("user").Hidden = true
	applyCommonFlagCompletion(filterCmd)
	rootCmd.AddCommand(filterCmd)
}
