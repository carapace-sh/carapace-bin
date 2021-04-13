package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove stopped containers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rmCmd).Standalone()

	rmCmd.Flags().BoolP("all", "a", false, "Deprecated - no effect.")
	rmCmd.Flags().BoolP("force", "f", false, "Don't ask to confirm removal")
	rmCmd.Flags().BoolP("stop", "s", false, "Stop the containers, if required, before removing")
	rmCmd.Flags().BoolS("v", "v", false, "Remove any anonymous volumes attached to containers")
	rootCmd.AddCommand(rmCmd)

	carapace.Gen(rmCmd).PositionalAnyCompletion(ActionServices())
}
