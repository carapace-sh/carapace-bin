package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var npm_whoamiCmd = &cobra.Command{
	Use:   "whoami",
	Short: "Display the name of the authenticated user",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(npm_whoamiCmd).Standalone()

	npm_whoamiCmd.Flags().Bool("publish", false, "Print username for the publish registry")
	npm_whoamiCmd.Flags().StringP("scope", "s", "", "Print username for the registry configured for a given scope")
	npmCmd.AddCommand(npm_whoamiCmd)

	// TODO scope
}
