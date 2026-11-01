package common

import (
	"github.com/spf13/cobra"
)

// AddJSONFlags adds JSON output flags
// (MixJSON in nix source).
func AddJSONFlags(cmd *cobra.Command) {
	cmd.Flags().Bool("json", false, "Produce output in JSON format, suitable for consumption by another program")
	cmd.Flags().Bool("no-pretty", false, "Print compact JSON output on a single line, even when the output is a terminal")
	cmd.Flags().Bool("pretty", false, "Print multi-line, indented JSON output for readability")
	cmd.MarkFlagsMutuallyExclusive("pretty", "no-pretty")
}
