package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dd",
	Short: "convert and copy a file",
	Long:  "https://keith.github.io/xcode-manpages/dd.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionMultiParts("=", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return carapace.ActionValuesDescribed(
					"bs", "Set both input and output block size to n bytes",
					"conv", "Convert the file as specified by the comma-separated option list",
					"count", "Copy only n input blocks",
					"files", "Set the number of input files to read from",
					"ibs", "Set the input block size to n bytes",
					"if", "Read input from file instead of the standard input",
					"iseq", "Set the sequence number of the input file",
					"obs", "Set the output block size to n bytes",
					"of", "Write output to file instead of the standard output",
					"oseq", "Set the sequence number of the output file",
					"seek", "Seek n blocks from the beginning of the output file before copying",
					"skip", "Skip n blocks from the beginning of the input file before copying",
				).Suffix("=")
			case 1:
				switch c.Parts[0] {
				case "if", "of":
					return carapace.ActionFiles()
				case "conv":
					return carapace.ActionValues(
						"ascii", "block", "ebcdic", "ibm", "lcase",
						"noerror", "notrunc", "osync", "swab", "ucase", "unblock",
					).UniqueList(",").NoSpace()
				default:
					return carapace.ActionValues()
				}
			default:
				return carapace.ActionValues()
			}
		}).NoSpace(),
	)
}
