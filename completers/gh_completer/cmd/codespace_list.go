package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var codespace_listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List your codespaces",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codespace_listCmd).Standalone()
	codespace_listCmd.Flags().StringP("jq", "q", "", "Filter JSON output using a jq `expression`")
	codespace_listCmd.Flags().StringSlice("json", []string{}, "Output JSON with the specified `fields`")
	codespace_listCmd.Flags().IntP("limit", "L", 30, "Maximum number of codespaces to list")
	codespace_listCmd.Flags().StringP("template", "t", "", "Format JSON output using a Go template")
	codespaceCmd.AddCommand(codespace_listCmd)

	carapace.Gen(codespace_listCmd).FlagCompletion(carapace.ActionMap{
		"json": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionCodespaceFields().Invoke(c).Filter(c.Parts).ToA()
		}),
	})
}
