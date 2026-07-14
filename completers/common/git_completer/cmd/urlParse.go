package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var urlParseCmd = &cobra.Command{
	Use:     "url-parse",
	Short:   "Parse and extract git URL components",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_low_level_helper].ID,
}

func init() {
	carapace.Gen(urlParseCmd).Standalone()

	urlParseCmd.Flags().StringP("component", "c", "", "which URL component to extract")
	rootCmd.AddCommand(urlParseCmd)

	carapace.Gen(urlParseCmd).FlagCompletion(carapace.ActionMap{
		"component": carapace.ActionValues("scheme", "user", "host", "port", "path"),
	})

	carapace.Gen(urlParseCmd).PositionalAnyCompletion(
		carapace.ActionValues(),
	)

	carapace.Gen(urlParseCmd).DashAnyCompletion(
		carapace.ActionPositional(urlParseCmd),
	)
}
