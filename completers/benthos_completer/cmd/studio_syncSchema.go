package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var studio_syncSchemaCmd = &cobra.Command{
	Use:   "sync-schema",
	Short: "Synchronizes the schema of this Benthos instance with a studio session",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(studio_syncSchemaCmd).Standalone()

	studio_syncSchemaCmd.Flags().StringP("session", "s", "", "The session ID to synchronize with.")
	studio_syncSchemaCmd.Flags().StringP("token", "t", "", "The single use token used to authenticate the request.")
	studioCmd.AddCommand(studio_syncSchemaCmd)
}
