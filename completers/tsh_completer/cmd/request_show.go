package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tsh"
	"github.com/spf13/cobra"
)

var request_showCmd = &cobra.Command{
	Use:     "show",
	Short:   "Show request details.",
	Aliases: []string{"details"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(request_showCmd).Standalone()

	request_showCmd.Flags().StringP("format", "f", "", "Format output (text, json, yaml)")
	requestCmd.AddCommand(request_showCmd)

	carapace.Gen(request_showCmd).FlagCompletion(carapace.ActionMap{
		"format": tsh.ActionFormats(),
	})
}
