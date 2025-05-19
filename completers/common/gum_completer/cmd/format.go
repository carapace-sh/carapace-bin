package cmd

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/chroma"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
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
	formatCmd.Flags().Bool("strip-ansi", false, "Strip ANSI sequences when reading from STDIN")
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
				if index := strings.LastIndex(c.Value, ":"); index != -1 {
					return gh.ActionEmojis().Prefix(c.Value[:index])
				}
			}
			return carapace.ActionValues()
		}),
	)
}
