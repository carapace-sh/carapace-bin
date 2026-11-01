package cmd

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dd",
	Short: "convert and copy a file",
	Long:  "https://linux.die.net/man/1/dd",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("help", false, "display this help and exit")
	rootCmd.Flags().Bool("version", false, "output version information and exit")

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionMultiParts("=", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				keys := make([]string, len(c.Args))
				for index, arg := range c.Args {
					keys[index] = strings.Split(arg, "=")[0]
				}
				return carapace.ActionValuesDescribed(
					"bs", "read and write up to BYTES bytes at a time",
					"cbs", "convert BYTES bytes at a time",
					"conv", "convert the file as per the comma separated symbol list",
					"count", "copy only N input blocks",
					"ibs", "read up to BYTES bytes at a time",
					"if", "read from FILE instead of stdin",
					"iflag", "read as per the comma separated symbol list",
					"obs", "write BYTES bytes at a time",
					"of", "write to FILE instead of stdout",
					"oflag", "write as per the comma separated symbol list",
					"seek", "skip N obs-sized blocks at start of output",
					"skip", "skip N ibs-sized blocks at start of input",
					"status", "The LEVEL of information to print to stderr",
				).Invoke(c).Filter(keys...).Suffix("=").ToA()
			case 1:
				switch c.Parts[0] {
				case "conv":
					return actionConvs().UniqueList(",")
				case "if":
					return carapace.ActionFiles()
				case "iflag":
					return actionIFlags().UniqueList(",")
				case "of":
					return carapace.ActionFiles()
				case "oflag":
					return actionOFlags().UniqueList(",")
				case "status":
					return carapace.ActionValues("none", "noxfer", "progress")
				default:
					return carapace.ActionValues()
				}
			default:
				return carapace.ActionValues()
			}
		}),
	)
}

func actionConvs() carapace.Action {
	return carapace.ActionValuesDescribed(
		"ascii", "from EBCDIC to ASCII",
		"block", "pad newline-terminated records with spaces to cbs-size",
		"ebcdic", "from ASCII to EBCDIC",
		"excl", "fail if the output file already exists",
		"fdatasync", "physically write output file data before finishing",
		"fsync", "likewise, but also write metadata",
		"ibm", "from ASCII to alternate EBCDIC",
		"lcase", "change upper case to lower case",
		"nocreat", "do not create the output file",
		"noerror", "continue after read errors",
		"notrunc", "do not truncate the output file",
		"sparse", "try to seek rather than write all-NUL output blocks",
		"swab", "swap every pair of input bytes",
		"sync", "pad every input block with NULs to ibs-size",
		"ucase", "change lower case to upper case",
		"unblock", "replace trailing spaces in cbs-size records with newline",
	)
}

func actionIFlags() carapace.Action {
	return carapace.ActionValuesDescribed(
		"append", "append mode",
		"count_bytes", "treat 'count=N' as a byte count",
		"direct", "use direct I/O for data",
		"directory", "fail unless a directory",
		"dsync", "use synchronized I/O for data",
		"fullblock", "accumulate full blocks of input",
		"noatime", "do not update access time",
		"nocache", "Request to drop cache",
		"noctty", "do not assign controlling terminal from file",
		"nofollow", "do not follow symlinks",
		"nonblock", "use non-blocking I/O",
		"skip_bytes", "treat 'skip=N' as a byte count",
		"sync", "likewise, but also for metadata",
	)
}

func actionOFlags() carapace.Action {
	return carapace.ActionValuesDescribed(
		"append", "append mode",
		"direct", "use direct I/O for data",
		"directory", "fail unless a directory",
		"dsync", "use synchronized I/O for data",
		"noatime", "do not update access time",
		"nocache", "Request to drop cache",
		"noctty", "do not assign controlling terminal from file",
		"nofollow", "do not follow symlinks",
		"nonblock", "use non-blocking I/O",
		"seek_bytes", "treat 'seek=N' as a byte count",
		"sync", "likewise, but also for metadata",
	)
}
