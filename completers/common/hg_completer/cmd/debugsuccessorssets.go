package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debugsuccessorssetsCmd = &cobra.Command{
	Use:    "debugsuccessorssets",
	Short:  "show set of successors for revision",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debugsuccessorssetsCmd).Standalone()

	debugsuccessorssetsCmd.Flags().Bool("closest", false, "return closest successors sets only")
	rootCmd.AddCommand(debugsuccessorssetsCmd)
}
