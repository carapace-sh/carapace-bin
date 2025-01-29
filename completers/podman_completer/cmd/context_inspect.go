package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var context_inspectCmd = &cobra.Command{
	Use:   "inspect [options] [CONTEXT] [CONTEXT...]",
	Short: "Inspect destination for a Podman service(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(context_inspectCmd).Standalone()

	context_inspectCmd.Flags().StringP("format", "f", "", "Custom Go template for printing connections")
	context_inspectCmd.Flags().BoolP("quiet", "q", false, "Custom Go template for printing connections")
	contextCmd.AddCommand(context_inspectCmd)
}
