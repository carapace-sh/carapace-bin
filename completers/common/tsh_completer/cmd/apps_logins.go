package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tsh"
	"github.com/spf13/cobra"
)

var apps_loginsCmd = &cobra.Command{
	Use:   "logins",
	Short: "List available logins for a Cloud console application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apps_loginsCmd).Standalone()

	apps_loginsCmd.Flags().StringP("format", "f", "text", "Format output (text, json, yaml).")
	appsCmd.AddCommand(apps_loginsCmd)

	carapace.Gen(apps_loginsCmd).FlagCompletion(carapace.ActionMap{
		"format": tsh.ActionFormats(),
	})

	carapace.Gen(apps_loginsCmd).PositionalCompletion(
		tsh.ActionApps(),
	)
}
