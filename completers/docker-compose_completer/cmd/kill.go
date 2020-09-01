package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var killCmd = &cobra.Command{
	Use:   "kill",
	Short: "Kill containers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(killCmd).Standalone()

	killCmd.Flags().StringS("s", "s", "", "SIGNAL to send to the container.")
	rootCmd.AddCommand(killCmd)

	carapace.Gen(killCmd).FlagCompletion(carapace.ActionMap{
		"s": carapace.ActionKillSignals(),
	})
}
