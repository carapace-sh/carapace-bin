package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-tool-asm",
	Short: "go assembler",
	Long:  "https://pkg.go.dev/cmd/asm",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringSliceS("D", "D", nil, "predefined symbol with optional simple value -D=identifier=value")
	rootCmd.Flags().StringSliceS("I", "I", nil, "include directory")
	rootCmd.Flags().BoolS("compiling-runtime", "compiling-runtime", false, "source to be compiled is part of the Go runtime")
	rootCmd.Flags().StringS("d", "d", "", "enable debugging settings")
	rootCmd.Flags().BoolS("debug", "debug", false, "dump instructions as they are parsed")
	rootCmd.Flags().BoolS("dynlink", "dynlink", false, "support references to Go symbols defined in other shared libraries")
	rootCmd.Flags().BoolS("gensymabis", "gensymabis", false, "write symbol ABI information to output file, don't assemble")
	rootCmd.Flags().BoolS("linkshared", "linkshared", false, "generate code that will be linked against Go shared libraries")
	rootCmd.Flags().StringS("o", "o", "", "output file")
	rootCmd.Flags().StringS("p", "p", "", "set expected package import to path")
	rootCmd.Flags().BoolS("shared", "shared", false, "generate code that can be linked into a shared library")
	rootCmd.Flags().StringS("spectre", "spectre", "", "enable spectre mitigations in list")
	rootCmd.Flags().StringS("trimpath", "trimpath", "", "remove prefix from recorded source file paths")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"I": carapace.ActionDirectories(),
		"d": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionMultiParts("=", func(c carapace.Context) carapace.Action {
				switch len(c.Parts) {
				case 0:
					return carapace.ActionValuesDescribed(
						"maymorestack", ",call named function before all stack growth checks",
						"pctab", "print named pc-value table",
					).Suffix("=")

				case 1:
					switch c.Parts[0] {
					case "pctab":
						return carapace.ActionValues("pctospadj", "pctofile", "pctoline", "pctoinline", "pctopcdata").NoSpace()
					default:
						return carapace.ActionValues()
					}

				default:
					return carapace.ActionValues()
				}
			})
		}),
		"o": carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
