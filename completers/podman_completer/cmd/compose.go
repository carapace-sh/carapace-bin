package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var composeCmd = &cobra.Command{
	Use:   "compose [options]",
	Short: "Run compose workloads via an external provider such as docker-compose or podman-compose",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(composeCmd).Standalone()

	rootCmd.AddCommand(composeCmd)
}
