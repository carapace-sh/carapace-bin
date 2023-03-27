package cmd

import (
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/chroma"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

var formatCmd = &cobra.Command{
	Use:   "format",
	Short: "Format a string using a template",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(formatCmd).Standalone()

	formatCmd.Flags().StringP("language", "l", "", "Programming language to parse code")
	formatCmd.Flags().String("theme", "", "Glamour theme to use for markdown formatting")
	formatCmd.Flags().StringP("type", "t", "", "Format to use (markdown,template,code,emoji)")
	rootCmd.AddCommand(formatCmd)

	carapace.Gen(formatCmd).FlagCompletion(carapace.ActionMap{
		"language": chroma.ActionLexers(),
		"theme": carapace.Batch(
			carapace.ActionValues("ascii", "dark", "dracula", "light", "notty", "pink").Tag("themes"),
			carapace.ActionFiles(),
		).ToA(),
		"type": carapace.ActionValues("markdown", "template", "code", "emoji"),
	})

	carapace.Gen(formatCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if formatCmd.Flag("type").Value.String() == "emoji" {
				if index := strings.LastIndex(c.CallbackValue, ":"); index != -1 {
					return gh.ActionEmojis().Prefix(c.CallbackValue[:index])
				}
			}
			return carapace.ActionValues()
		}),
	)
}
