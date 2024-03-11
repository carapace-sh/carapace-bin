package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tsh"
	"github.com/spf13/cobra"
)

var request_lsCmd = &cobra.Command{
	Use:     "ls",
	Short:   "List access requests.",
	Aliases: []string{"list"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(request_lsCmd).Standalone()

	request_lsCmd.Flags().StringP("format", "f", "", "Format output (text, json, yaml)")
	request_lsCmd.Flags().Bool("my-requests", false, "Only show requests created by current user")
	request_lsCmd.Flags().Bool("no-my-requests", false, "Only show requests created by current user")
	request_lsCmd.Flags().Bool("no-reviewable", false, "Only show requests reviewable by current user")
	request_lsCmd.Flags().Bool("no-suggested", false, "Only show requests that suggest current user as reviewer")
	request_lsCmd.Flags().Bool("reviewable", false, "Only show requests reviewable by current user")
	request_lsCmd.Flags().Bool("suggested", false, "Only show requests that suggest current user as reviewer")
	request_lsCmd.Flag("no-my-requests").Hidden = true
	request_lsCmd.Flag("no-reviewable").Hidden = true
	request_lsCmd.Flag("no-suggested").Hidden = true
	requestCmd.AddCommand(request_lsCmd)

	carapace.Gen(request_lsCmd).FlagCompletion(carapace.ActionMap{
		"format": tsh.ActionFormats(),
	})
}
