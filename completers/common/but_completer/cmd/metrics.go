package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var metricsCmd = &cobra.Command{
	Use:    "metrics",
	Short:  "INTERNAL: If metrics are permitted, this subcommand handles posthog event creation",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(metricsCmd).Standalone()

	metricsCmd.Flags().String("command-name", "", "")
	metricsCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	metricsCmd.Flags().String("props", "", "")
	metricsCmd.MarkFlagRequired("command-name")
	metricsCmd.MarkFlagRequired("props")
	rootCmd.AddCommand(metricsCmd)
}