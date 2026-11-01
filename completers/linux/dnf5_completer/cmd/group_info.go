package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var groupInfoCmd = &cobra.Command{
	Use:   "info [options] [<group-spec>...]",
	Short: "print details about comps groups",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(groupInfoCmd).Standalone()

	groupInfoCmd.Flags().Bool("available", false, "List available groups")
	groupInfoCmd.Flags().String("contains-pkgs", "", "Filter by packages in group")
	groupInfoCmd.Flags().Bool("hidden", false, "Include hidden groups")
	groupInfoCmd.Flags().Bool("installed", false, "List installed groups")

	groupCmd.AddCommand(groupInfoCmd)
}
