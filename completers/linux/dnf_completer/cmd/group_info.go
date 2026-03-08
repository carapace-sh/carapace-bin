package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var groupInfoCmd = &cobra.Command{
	Use:   "info [group-spec]...",
	Short: "print detailed information about groups",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(groupInfoCmd).Standalone()

	groupInfoCmd.Flags().Bool("available", false, "show only available groups")
	groupInfoCmd.Flags().String("contains-pkgs", "", "show only groups containing packages")
	groupInfoCmd.Flags().Bool("hidden", false, "show also hidden groups")
	groupInfoCmd.Flags().Bool("installed", false, "show only installed groups")

	groupCmd.AddCommand(groupInfoCmd)
}
