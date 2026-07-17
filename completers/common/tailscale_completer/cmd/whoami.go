package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var whoamiCmd = &cobra.Command{
	Use:   "whoami",
	Short: "Show the machine and user identity of the current machine",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(whoamiCmd).Standalone()

	whoamiCmd.Flags().Bool("json", false, "output in JSON format")
	rootCmd.AddCommand(whoamiCmd)
}
