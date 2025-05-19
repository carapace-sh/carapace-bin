package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/env"
	"github.com/spf13/cobra"
)

var kustomizeCmd = &cobra.Command{
	Use:     "kustomize DIR",
	Short:   "Build a kustomization target from a directory or URL",
	GroupID: "advanced",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kustomizeCmd).Standalone()

	kustomizeCmd.Flags().Bool("as-current-user", false, "use the uid and gid of the command executor to run the function in the container")
	kustomizeCmd.Flags().Bool("enable-alpha-plugins", false, "enable kustomize plugins")
	kustomizeCmd.Flags().Bool("enable-helm", false, "Enable use of the Helm chart inflator generator.")
	kustomizeCmd.Flags().Bool("enable-managedby-label", false, "enable adding app.kubernetes.io/managed-by")
	kustomizeCmd.Flags().StringSliceP("env", "e", nil, "a list of environment variables to be used by functions")
	kustomizeCmd.Flags().StringSlice("helm-api-versions", nil, "Kubernetes api versions used by Helm for Capabilities.APIVersions")
	kustomizeCmd.Flags().String("helm-command", "", "helm command (path to executable)")
	kustomizeCmd.Flags().Bool("helm-debug", false, "Enable debug output from the Helm chart inflator generator.")
	kustomizeCmd.Flags().String("helm-kube-version", "", "Kubernetes version used by Helm for Capabilities.KubeVersion")
	kustomizeCmd.Flags().String("load-restrictor", "", "if set to 'LoadRestrictionsNone', local kustomizations may load files from outside their root. This does, however, break the relocatability of the kustomization.")
	kustomizeCmd.Flags().StringSlice("mount", nil, "a list of storage options read from the filesystem")
	kustomizeCmd.Flags().Bool("network", false, "enable network access for functions that declare it")
	kustomizeCmd.Flags().String("network-name", "", "the docker network to run the container in")
	kustomizeCmd.Flags().StringP("output", "o", "", "If specified, write output to this path.")
	kustomizeCmd.Flags().String("reorder", "", "Reorder the resources just before output. Use 'legacy' to apply a legacy reordering (Namespaces first, Webhooks last, etc). Use 'none' to suppress a final reordering.")
	kustomizeCmd.Flag("enable-managedby-label").Hidden = true
	kustomizeCmd.Flag("reorder").Hidden = true
	rootCmd.AddCommand(kustomizeCmd)

	carapace.Gen(kustomizeCmd).FlagCompletion(carapace.ActionMap{
		"env":     env.ActionNameValues(false),
		"output":  carapace.ActionFiles(),
		"reorder": carapace.ActionValues("legacy", "none"),
	})

	carapace.Gen(kustomizeCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
