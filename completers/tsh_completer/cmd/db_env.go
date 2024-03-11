package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tsh"
	"github.com/spf13/cobra"
)

var db_envCmd = &cobra.Command{
	Use:   "env",
	Short: "Print environment variables for the configured database.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(db_envCmd).Standalone()

	db_envCmd.Flags().String("db", "", "Print environment for the specified database.")
	db_envCmd.Flags().StringP("format", "f", "", "Format output (text, json, yaml)")
	db_envCmd.Flag("db").Hidden = true
	dbCmd.AddCommand(db_envCmd)

	carapace.Gen(db_envCmd).FlagCompletion(carapace.ActionMap{
		"format": tsh.ActionFormats(),
	})
}
