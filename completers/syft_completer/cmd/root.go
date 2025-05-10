package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/syft"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "syft",
	Short: "Generate a package SBOM",
	Long:  "https://github.com/anchore/syft",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.PersistentFlags().StringP("config", "c", "", "application config file")
	rootCmd.PersistentFlags().BoolP("quiet", "q", false, "suppress all logging output")
	rootCmd.PersistentFlags().CountP("verbose", "v", "increase verbosity (-v = info, -vv = debug)")
	addCommonFlags(rootCmd)

	// TODO platform, catalogers
	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"config": carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		syft.ActionSources(),
	)
}

func addCommonFlags(cmd *cobra.Command) {
	cmd.Flags().StringArray("catalogers", nil, "enable one or more package catalogers")
	cmd.Flags().StringArray("exclude", nil, "exclude paths from being scanned using a glob expression")
	cmd.Flags().String("file", "", "file to write the default report output to (default is STDOUT)")
	cmd.Flags().String("name", "", "set the name of the target being analyzed")
	cmd.Flags().StringArrayP("output", "o", []string{"syft-table"}, "report output format, options=[syft-json cyclonedx-xml cyclonedx-json github-json spdx-tag-value spdx-json syft-table syft-text template]")
	cmd.Flags().String("platform", "", "an optional platform specifier for container image sources (e.g. 'linux/arm64', 'linux/arm64/v8', 'arm64', 'linux')")
	cmd.Flags().StringP("scope", "s", "Squashed", "selection of layers to catalog, options=[Squashed AllLayers]")
	cmd.Flags().StringP("template", "t", "", "specify the path to a Go template file")

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"file": carapace.ActionFiles(),
		"output": carapace.ActionMultiParts("=", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return syft.ActionOutputFormats()

			case 1:
				return carapace.ActionFiles()

			default:
				return carapace.ActionValues()
			}
		}),
		"scope":    carapace.ActionValues("Squashed", "AllLayers"),
		"template": carapace.ActionFiles(),
	})
}
