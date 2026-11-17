package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var whoamiCmd = &cobra.Command{
	Use:   "whoami",
	Short: "Show your identity, server, and Rex location",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(whoamiCmd).Standalone()

	whoamiCmd.Flags().Bool("json", false, "output identity, server, and location as JSON")

	rootCmd.AddCommand(whoamiCmd)
}
