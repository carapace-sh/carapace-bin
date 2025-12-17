package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/kubectl"
	"github.com/spf13/cobra"
)

var clusterInfo_dumpCmd = &cobra.Command{
	Use:   "dump",
	Short: "Dump relevant information for debugging and diagnosis",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clusterInfo_dumpCmd).Standalone()

	clusterInfo_dumpCmd.Flags().BoolP("all-namespaces", "A", false, "If true, dump all namespaces.  If true, --namespaces is ignored.")
	clusterInfo_dumpCmd.Flags().Bool("allow-missing-template-keys", false, "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.")
	clusterInfo_dumpCmd.Flags().StringSlice("namespaces", nil, "A comma separated list of namespaces to dump.")
	clusterInfo_dumpCmd.Flags().StringP("output", "o", "", "Output format. One of: (json, yaml, kyaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).")
	clusterInfo_dumpCmd.Flags().String("output-directory", "", "Where to output the files.  If empty or '-' uses stdout, otherwise creates a directory hierarchy in that directory")
	clusterInfo_dumpCmd.Flags().String("pod-running-timeout", "", "The length of time (like 5s, 2m, or 3h, higher than zero) to wait until at least one pod is running")
	clusterInfo_dumpCmd.Flags().Bool("show-managed-fields", false, "If true, keep the managedFields when printing objects in JSON or YAML format.")
	clusterInfo_dumpCmd.Flags().String("template", "", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].")
	clusterInfoCmd.AddCommand(clusterInfo_dumpCmd)

	carapace.Gen(clusterInfo_dumpCmd).FlagCompletion(carapace.ActionMap{
		"namespaces": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			var namespace string
			if !clusterInfoCmd.Flag("all-namespaces").Changed {
				namespace = rootCmd.Flag("namespace").Value.String()
			}
			return kubectl.ActionResources(kubectl.ResourceOpts{
				Context:   rootCmd.Flag("context").Value.String(),
				Namespace: namespace,
				Types:     "namespaces",
			})
		}),
		"output":           kubectl.ActionOutputFormats(),
		"output-directory": carapace.ActionDirectories(),
		"template":         carapace.ActionFiles(),
	})
}
