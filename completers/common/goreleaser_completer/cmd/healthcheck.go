package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var healthcheckCmd = &cobra.Command{
	Use:     "healthcheck",
	Short:   "Checks if needed tools are installed",
	Aliases: []string{"hc"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(healthcheckCmd).Standalone()

	healthcheckCmd.Flags().StringP("config", "f", "", "Configuration file")
	healthcheckCmd.Flags().BoolP("quiet", "q", false, "Quiet mode: no output")
	rootCmd.AddCommand(healthcheckCmd)

	carapace.Gen(healthcheckCmd).FlagCompletion(carapace.ActionMap{
		"config": carapace.ActionFiles(),
	})
}
