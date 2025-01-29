package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var machine_listCmd = &cobra.Command{
	Use:     "list [options]",
	Short:   "List machines",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(machine_listCmd).Standalone()

	machine_listCmd.Flags().Bool("all-providers", false, "Show machines from all providers")
	machine_listCmd.Flags().String("format", "", "Format volume output using JSON or a Go template")
	machine_listCmd.Flags().BoolP("noheading", "n", false, "Do not print headers")
	machine_listCmd.Flags().BoolP("quiet", "q", false, "Show only machine names")
	machineCmd.AddCommand(machine_listCmd)
}
