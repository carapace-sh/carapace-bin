package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var planCmd = &cobra.Command{
	Use:   "plan",
	Short: "Preview the plan used to build your environment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(planCmd).Standalone()
	planCmd.Flags().StringP("config", "c", "", "path to directory containing a devbox.json config file")
	rootCmd.AddCommand(planCmd)

	carapace.Gen(planCmd).FlagCompletion(carapace.ActionMap{
		"config": carapace.ActionDirectories(),
	})

	carapace.Gen(planCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
