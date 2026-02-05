package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var config_metricsCmd = &cobra.Command{
	Use:   "metrics",
	Short: "View or set metrics collection",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_metricsCmd).Standalone()

	config_metricsCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	configCmd.AddCommand(config_metricsCmd)

	carapace.Gen(config_metricsCmd).PositionalCompletion(
		carapace.ActionValues("enable", "disable").StyleF(style.ForKeyword),
	)
}
