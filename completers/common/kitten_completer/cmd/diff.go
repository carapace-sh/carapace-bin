package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var diffCmd = &cobra.Command{
	Use:   "diff",
	Short: "Pretty, side-by-side diffing of files and images",
}

func init() {
	rootCmd.AddCommand(diffCmd)
	carapace.Gen(diffCmd).Standalone()

	diffCmd.Flags().StringArray("config", nil, "Specify a path to the configuration file(s) to use. All configuration files are merged onto the builtin diff.conf, overriding the builtin values. This option can be specified multiple times to read multiple configuration files in sequence, which are merged. Use the special value NONE to not load any config file.")
	diffCmd.Flags().Int("context", 0, "Number of lines of context to show between changes. Negative values use the number set in diff.conf.")
	diffCmd.Flags().BoolP("help", "h", false, "Show help for this command")
	diffCmd.Flags().StringP("override", "o", "", "Override individual configuration options, can be specified multiple times. Syntax: name=value. For example: -o background=gray")

	carapace.Gen(diffCmd).FlagCompletion(carapace.ActionMap{
		"config": carapace.Batch(carapace.ActionFiles(), carapace.ActionValues("NONE", "-", "/dev/stdin")).ToA(),
	})

	carapace.Gen(diffCmd).PositionalAnyCompletion(carapace.Batch(
		carapace.ActionFiles(),
		carapace.ActionDirectories(),
	).ToA())
}
