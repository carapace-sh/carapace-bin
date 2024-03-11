package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tsh"
	"github.com/spf13/cobra"
)

var recordings_lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List recorded sessions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(recordings_lsCmd).Standalone()

	recordings_lsCmd.Flags().StringP("format", "f", "", "Format output (text, json, yaml). Defaults to 'text'.")
	recordings_lsCmd.Flags().String("from-utc", "", "Start of time range in which recordings are listed. Format 2006-01-02. Defaults to 24 hours ago.")
	recordings_lsCmd.Flags().String("last", "", "Duration into the past from which session recordings should be listed. Format 5h30m40s")
	recordings_lsCmd.Flags().String("limit", "", "Maximum number of recordings to show. Default 50.")
	recordings_lsCmd.Flags().String("to-utc", "", "End of time range in which recordings are listed. Format 2006-01-02. Defaults to current time.")
	recordingsCmd.AddCommand(recordings_lsCmd)

	carapace.Gen(recordings_lsCmd).FlagCompletion(carapace.ActionMap{
		"format": tsh.ActionFormats(),
	})
}
