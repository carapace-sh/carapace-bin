package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/syft"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "grype [IMAGE]",
	Short: "A vulnerability scanner for container images, filesystems, and SBOMs",
	Long:  "https://github.com/anchore/grype",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute(opts ...func(cmd *cobra.Command)) error {
	for _, opt := range opts {
		opt(rootCmd)
	}
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("add-cpes-if-none", false, "generate CPEs for packages with no CPE data")
	rootCmd.Flags().Bool("by-cve", false, "orient results by CVE instead of the original vulnerability ID when possible")
	rootCmd.PersistentFlags().StringP("config", "c", "", "application config file")
	rootCmd.Flags().String("distro", "", "distro to match against in the format: <distro>:<version>")
	rootCmd.Flags().StringArray("exclude", nil, "exclude paths from being scanned using a glob expression")
	rootCmd.Flags().StringP("fail-on", "f", "", "set the return code to 1 if a vulnerability is found with a severity >= the given severity")
	rootCmd.Flags().String("file", "", "file to write the report output to (default is STDOUT)")
	rootCmd.Flags().Bool("only-fixed", false, "ignore matches for vulnerabilities that are not fixed")
	rootCmd.Flags().Bool("only-notfixed", false, "ignore matches for vulnerabilities that are fixed")
	rootCmd.Flags().StringP("output", "o", "", "report output formatter")
	rootCmd.Flags().String("platform", "", "an optional platform specifier for container image sources (e.g. 'linux/arm64', 'linux/arm64/v8', 'arm64', 'linux')")
	rootCmd.PersistentFlags().BoolP("quiet", "q", false, "suppress all logging output")
	rootCmd.Flags().StringP("scope", "s", "Squashed", "selection of layers to analyze")
	rootCmd.Flags().Bool("show-suppressed", false, "show suppressed/ignored vulnerabilities in the output (only supported with table output format)")
	rootCmd.Flags().StringP("template", "t", "", "specify the path to a Go template file (requires 'template' output to be selected)")
	rootCmd.PersistentFlags().CountP("verbose", "v", "increase verbosity")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"config":   carapace.ActionFiles(),
		"fail-on":  carapace.ActionValues("negligible", "low", "medium", "high", "critical"),
		"file":     carapace.ActionFiles(),
		"output":   syft.ActionOutputFormats(),
		"scope":    carapace.ActionValues("Squashed", "AllLayers"),
		"template": carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		syft.ActionSources(),
	)
}
