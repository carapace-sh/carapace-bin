package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/ng_completer/cmd/action"
	"github.com/spf13/cobra"
)

var i18nExtractCmd = &cobra.Command{
	Use:   "i18n-extract",
	Short: "Extracts i18n messages from source code",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(i18nExtractCmd).Standalone()

	i18nExtractCmd.Flags().String("browser-target", "", "A browser builder target to extract i18n messages")
	i18nExtractCmd.Flags().StringP("configuration", "c", "", "One or more named builder configurations")
	i18nExtractCmd.Flags().String("format", "", "Output format for the generated file.")
	i18nExtractCmd.Flags().String("out-file", "", "Name of the file to output.")
	i18nExtractCmd.Flags().String("output-path", "", "Path where output will be placed.")
	i18nExtractCmd.Flags().Bool("progress", false, "Log progress to the console.")
	rootCmd.AddCommand(i18nExtractCmd)

	carapace.Gen(i18nExtractCmd).FlagCompletion(carapace.ActionMap{
		"configuration": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			if len(c.Args) < 1 {
				return carapace.ActionValues()
			}
			return action.ActionBuilderConfigurations(c.Args[0]).Invoke(c).Filter(c.Parts).ToA()
		}),
		"format":      carapace.ActionValues("xmb", "xlf", "xlif", "xliff", "xlf2", "xliff2", "json", "arb", "legacy-migrate"),
		"out-file":    carapace.ActionFiles(),
		"output-path": carapace.ActionDirectories(),
	})

	carapace.Gen(i18nExtractCmd).PositionalCompletion(
		action.ActionProjects(),
	)
}
