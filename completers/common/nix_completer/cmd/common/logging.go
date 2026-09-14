package common

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

// AddLoggingFlags adds logging verbosity flags
// (MixCommonArgs in nix source).
func AddLoggingFlags(cmd *cobra.Command) {
	cmd.Flags().Bool("debug", false, "Set the logging verbosity level to 'debug'")
	cmd.Flags().String("log-format", "", "Set the format of log output")
	cmd.Flags().BoolP("print-build-logs", "L", false, "Print full build logs on standard error")
	cmd.Flags().Bool("quiet", false, "Decrease the logging verbosity level")
	cmd.Flags().BoolP("verbose", "v", false, "Increase the logging verbosity level")

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"log-format": carapace.ActionValues("raw", "internal-json", "bar", "bar-with-logs"),
	})
}
