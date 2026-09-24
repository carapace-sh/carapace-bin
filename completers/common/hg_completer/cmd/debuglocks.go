package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debuglocksCmd = &cobra.Command{
	Use:    "debuglocks",
	Short:  "show or modify state of locks",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debuglocksCmd).Standalone()

	debuglocksCmd.Flags().BoolP("force-free-lock", "L", false, "free the store lock (DANGEROUS)")
	debuglocksCmd.Flags().BoolP("force-free-wlock", "W", false, "free the working state lock (DANGEROUS)")
	debuglocksCmd.Flags().BoolP("set-lock", "s", false, "set the store lock until stopped")
	debuglocksCmd.Flags().BoolP("set-wlock", "S", false, "set the working state lock until stopped")
	rootCmd.AddCommand(debuglocksCmd)
}
