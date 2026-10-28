package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chart_exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Render a Helm chart and export each resource as a separate YAML file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chart_exportCmd).Standalone()

	chart_exportCmd.Flags().Bool("crds", true, "Include CRDs in the output")
	chart_exportCmd.Flags().StringP("namespace", "n", "", "Kubernetes namespace for rendered manifests")
	chart_exportCmd.Flags().StringP("output", "o", "./yaml", "Output directory for rendered YAML files")
	chart_exportCmd.Flags().String("release", "openfaas", "Helm release name")
	chart_exportCmd.Flags().StringArray("set", nil, "Set individual values (key=value)")
	chart_exportCmd.Flags().StringArray("values", nil, "Path to values file(s) to use during rendering")
	chartCmd.AddCommand(chart_exportCmd)

	carapace.Gen(chart_exportCmd).FlagCompletion(carapace.ActionMap{
		"output": carapace.ActionDirectories(),
		"values": carapace.ActionFiles("yaml", "yml"),
	})

	carapace.Gen(chart_exportCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
