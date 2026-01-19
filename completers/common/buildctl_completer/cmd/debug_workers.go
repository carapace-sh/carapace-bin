package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debug_workersCmd = &cobra.Command{
	Use:   "workers",
	Short: "list workers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_workersCmd).Standalone()

	debug_workersCmd.Flags().StringP("filter", "f", "", "containerd-style filter string slice")
	debug_workersCmd.Flags().String("format", "", "Format the output using the given Go template")
	debug_workersCmd.Flags().BoolP("verbose", "v", false, "Verbose output")
	debugCmd.AddCommand(debug_workersCmd)
}
