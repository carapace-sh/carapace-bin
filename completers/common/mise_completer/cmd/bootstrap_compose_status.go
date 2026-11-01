package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bootstrap_compose_statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show the state of Docker Compose projects from `[bootstrap.compose]`",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootstrap_compose_statusCmd).Standalone()

	bootstrap_compose_statusCmd.Flags().BoolP("help", "h", false, "Print help")
	bootstrap_compose_statusCmd.Flags().BoolP("json", "J", false, "Output in JSON format")
	bootstrap_compose_statusCmd.Flags().Bool("missing", false, "Exit with status 1 if any Compose project differs from the config (all are still listed)")
	bootstrap_composeCmd.AddCommand(bootstrap_compose_statusCmd)
}
