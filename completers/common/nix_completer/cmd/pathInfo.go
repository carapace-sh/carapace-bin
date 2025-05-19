package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/nix"
	"github.com/spf13/cobra"
)

var pathInfoCmd = &cobra.Command{
	Use:     "path-info",
	Short:   "query information about store paths",
	GroupID: "infrequently used",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pathInfoCmd).Standalone()

	pathInfoCmd.Flags().BoolP("closure-size", "S", false, "Print the sum of the sizes of the NAR serialisations of the closure of each path")
	pathInfoCmd.Flags().BoolP("human-readable", "h", false, "With -s and -S, print sizes in a human-friendly format such as 5.67G")
	pathInfoCmd.Flags().Bool("json", false, "Produce output in JSON format, suitable for consumption by another program")
	pathInfoCmd.Flags().Bool("sigs", false, "Show signatures")
	pathInfoCmd.Flags().BoolP("size", "s", false, "Print the size of the NAR serialisation of each path")
	rootCmd.AddCommand(pathInfoCmd)

	addEvaluationFlags(pathInfoCmd)
	addFlakeFlags(pathInfoCmd)
	addLoggingFlags(pathInfoCmd)

	carapace.Gen(pathInfoCmd).FlagCompletion(carapace.ActionMap{
		"inputs-from": carapace.Batch(
			carapace.ActionDirectories(),
			nix.ActionFlakes(),
		).ToA(),
		"output-lock-file":    carapace.ActionFiles(),
		"reference-lock-file": carapace.ActionFiles("lock"),
	})
	carapace.Gen(pathInfoCmd).PositionalAnyCompletion(nix.ActionInstallables())
}
