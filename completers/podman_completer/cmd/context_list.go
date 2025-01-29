package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var context_listCmd = &cobra.Command{
	Use:     "list [options]",
	Short:   "List destination for the Podman service(s)",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(context_listCmd).Standalone()

	context_listCmd.Flags().StringP("format", "f", "", "Custom Go template for printing connections")
	context_listCmd.Flags().BoolP("quiet", "q", false, "Custom Go template for printing connections")
	contextCmd.AddCommand(context_listCmd)
}
