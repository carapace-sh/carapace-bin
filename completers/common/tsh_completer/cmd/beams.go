package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var beamsCmd = &cobra.Command{
	Use:     "beams",
	Short:   "View, manage and run beams. Beams are ephemeral, sandbox VMs built for agentic workloads.",
	Aliases: []string{"beam"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(beamsCmd).Standalone()

	rootCmd.AddCommand(beamsCmd)
}
