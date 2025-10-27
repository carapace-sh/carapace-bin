package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/minikube_completer/cmd/action"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:     "status",
	Short:   "Gets the status of a local Kubernetes cluster",
	GroupID: "basic",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(statusCmd).Standalone()

	statusCmd.Flags().StringP("format", "f", "", "Go template format string for the status output.  The format for Go templates can be found here: https://pkg.go.dev/text/template")
	statusCmd.Flags().StringP("layout", "l", "", "output layout (EXPERIMENTAL, JSON only): 'nodes' or 'cluster'")
	statusCmd.Flags().StringP("node", "n", "", "The node to check status for. Defaults to control plane. Leave blank with default format for status on all nodes.")
	statusCmd.Flags().StringP("output", "o", "", "minikube status --output OUTPUT. json, text")
	statusCmd.Flags().StringP("watch", "w", "", "Continuously listing/getting the status with optional interval duration.")
	statusCmd.Flag("watch").NoOptDefVal = " "
	rootCmd.AddCommand(statusCmd)

	carapace.Gen(statusCmd).FlagCompletion(carapace.ActionMap{
		"node":   action.ActionNodes(),
		"output": carapace.ActionValues("json", "text"),
	})
}
